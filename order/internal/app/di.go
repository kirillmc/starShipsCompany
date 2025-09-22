package app

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	kafkaConverter "github.com/kirillmc/starShipsCompany/order/internal/converter/kafka"
	"github.com/kirillmc/starShipsCompany/order/internal/converter/kafka/decoder"
	"github.com/kirillmc/starShipsCompany/order/internal/service/consumer/orderConsumer"
	"github.com/kirillmc/starShipsCompany/order/internal/service/producer/orderProducer"
	wrappedKafka "github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/kirillmc/starShipsCompany/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/kirillmc/starShipsCompany/platform/pkg/kafka/producer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	kafkaMiddleware "github.com/kirillmc/starShipsCompany/platform/pkg/middleware/kafka"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	orderAPI "github.com/kirillmc/starShipsCompany/order/internal/api/order/v1"
	grpcClients "github.com/kirillmc/starShipsCompany/order/internal/client/grpc"
	inventoryClient "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/inventory/v1"
	paymentClient "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/payment/v1"
	"github.com/kirillmc/starShipsCompany/order/internal/config"
	"github.com/kirillmc/starShipsCompany/order/internal/migrator"
	repo "github.com/kirillmc/starShipsCompany/order/internal/repository/pg"
	orderRepository "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/order"
	"github.com/kirillmc/starShipsCompany/order/internal/service"
	orderService "github.com/kirillmc/starShipsCompany/order/internal/service/order"
	"github.com/kirillmc/starShipsCompany/platform/pkg/closer"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type diContainer struct {
	orderV1Handler orderV1.Handler

	orderService service.Service
	orderRepo    repo.OrderRepository
	pgxPool      *pgxpool.Pool

	inventoryClient grpcClients.InventoryClient
	paymentClient   grpcClients.PaymentClient

	orderAssembledConsumerService service.ConsumerService
	orderAssembledConsumerGroup   sarama.ConsumerGroup
	orderAssembledConsumer        wrappedKafka.Consumer
	orderAssembledDecoder         kafkaConverter.Decoder

	orderPaidProducerService service.OrderProducerService
	orderPaidProducer        wrappedKafka.Producer
	syncProducer             sarama.SyncProducer
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderV1Handler(ctx context.Context) orderV1.Handler {
	if d.orderV1Handler == nil {
		d.orderV1Handler = orderAPI.NewAPI(d.OrderService(ctx))
	}

	return d.orderV1Handler
}

func (d *diContainer) OrderService(ctx context.Context) service.Service {
	if d.orderService == nil {
		d.orderService = orderService.NewService(
			d.PgxPool(ctx),
			d.InventoryClient(),
			d.PaymentClient(),
			d.OrderPaidProducerService(),
			d.OrderRepository(ctx),
		)
	}

	return d.orderService
}

func (d *diContainer) OrderRepository(ctx context.Context) repo.OrderRepository {
	if d.orderRepo == nil {
		d.orderRepo = orderRepository.NewRepository(d.PgxPool(ctx))
	}

	return d.orderRepo
}

func (d *diContainer) PgxPool(ctx context.Context) *pgxpool.Pool {
	if d.pgxPool == nil {
		var err error
		d.pgxPool, err = pgxpool.New(ctx, config.AppConfig().Postgres.URI())
		if err != nil {
			panic(fmt.Sprintf("failed to connect to database: %s\n", err))
		}

		closer.AddNamed("Pgx pool", func(ctx context.Context) error {
			d.pgxPool.Close()
			return nil
		})

		err = d.pgxPool.Ping(ctx)
		if err != nil {
			panic(fmt.Sprintf("db is not available: %s\n", err))
		}

		migratorRunner := migrator.NewMigrator(stdlib.OpenDB(*d.pgxPool.Config().Copy().ConnConfig),
			config.AppConfig().Postgres.MigrationsDir())

		err = migratorRunner.Up()
		if err != nil {
			panic(fmt.Sprintf("failed to migrate db: %s\n", err))
		}
	}

	return d.pgxPool
}

func (d *diContainer) InventoryClient() grpcClients.InventoryClient {
	if d.inventoryClient == nil {
		connInventory, err := grpc.NewClient(
			config.AppConfig().ExtDep.InventoryAddress(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to connect to inventory service: %s\n", err))
		}

		closer.AddNamed("Inventory client", func(ctx context.Context) error { return connInventory.Close() })

		d.inventoryClient = inventoryClient.NewClient(inventoryV1.NewInventoryServiceClient(connInventory))
	}

	return d.inventoryClient
}

func (d *diContainer) PaymentClient() grpcClients.PaymentClient {
	if d.paymentClient == nil {
		connPayment, err := grpc.NewClient(
			config.AppConfig().ExtDep.PaymentAddress(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to connect to payment service: %s\n", err))
		}

		closer.AddNamed("Payment client", func(ctx context.Context) error { return connPayment.Close() })

		d.paymentClient = paymentClient.NewClient(paymentV1.NewPaymentServiceClient(connPayment))
	}

	return d.paymentClient
}

func (d *diContainer) OrderAssembledConsumerService(ctx context.Context) service.ConsumerService {
	if d.orderAssembledConsumerService == nil {
		d.orderAssembledConsumerService = orderConsumer.NewService(
			d.OrderAssembledConsumer(),
			d.OrderAssembledDecoder(),
			d.OrderService(ctx),
		)
	}

	return d.orderAssembledConsumerService
}

func (d *diContainer) OrderAssembledConsumer() wrappedKafka.Consumer {
	if d.orderAssembledConsumer == nil {
		d.orderAssembledConsumer = wrappedKafkaConsumer.NewConsumer(
			d.OrderAssembledConsumerGroup(),
			[]string{
				config.AppConfig().OrderAssembledConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}

	return d.orderAssembledConsumer
}

func (d *diContainer) OrderAssembledConsumerGroup() sarama.ConsumerGroup {
	if d.orderAssembledConsumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssembledConsumer.GroupID(),
			config.AppConfig().OrderAssembledConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create order assembled consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka order assembled consumer group", func(ctx context.Context) error {
			return d.orderAssembledConsumerGroup.Close()
		})

		d.orderAssembledConsumerGroup = consumerGroup
	}

	return d.orderAssembledConsumerGroup
}

func (d *diContainer) OrderAssembledDecoder() kafkaConverter.Decoder {
	if d.orderAssembledDecoder == nil {
		d.orderAssembledDecoder = decoder.NewOrderAssembledDecoder()
	}

	return d.orderAssembledDecoder
}

func (d *diContainer) OrderPaidProducerService() service.OrderProducerService {
	if d.orderPaidProducerService == nil {
		d.orderPaidProducerService = orderProducer.NewService(d.OrderPaidProducer())
	}

	return d.orderPaidProducerService
}

func (d *diContainer) OrderPaidProducer() wrappedKafka.Producer {
	if d.orderPaidProducer == nil {
		d.orderPaidProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().OrderPaidProducer.Topic(),
			logger.Logger(),
		)
	}

	return d.orderPaidProducer
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidProducer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create sync producer: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka sync producer", func(ctx context.Context) error {
			return p.Close()
		})

		d.syncProducer = p
	}

	return d.syncProducer
}

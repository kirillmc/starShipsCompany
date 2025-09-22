package app

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/kirillmc/starShipsCompany/assembly/internal/config"
	kafkaConverter "github.com/kirillmc/starShipsCompany/assembly/internal/converter/kafka"
	"github.com/kirillmc/starShipsCompany/assembly/internal/converter/kafka/decoder"
	"github.com/kirillmc/starShipsCompany/assembly/internal/service"
	"github.com/kirillmc/starShipsCompany/assembly/internal/service/consumer/orderConsumer"
	"github.com/kirillmc/starShipsCompany/assembly/internal/service/producer/orderProducer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/closer"
	wrappedKafka "github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	wrappedKafkaConsumer "github.com/kirillmc/starShipsCompany/platform/pkg/kafka/consumer"
	wrappedKafkaProducer "github.com/kirillmc/starShipsCompany/platform/pkg/kafka/producer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	kafkaMiddleware "github.com/kirillmc/starShipsCompany/platform/pkg/middleware/kafka"
)

type diContainer struct {
	orderAssembledProducerService service.OrderProducerService
	orderPaidConsumerService      service.ConsumerService

	consumerGroup     sarama.ConsumerGroup
	orderPaidConsumer wrappedKafka.Consumer

	orderPaidDecoder       kafkaConverter.OrderPaidDecoder
	syncProducer           sarama.SyncProducer
	orderAssembledProducer wrappedKafka.Producer
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) OrderPaidConsumerService() service.ConsumerService {
	if d.orderPaidConsumerService == nil {
		d.orderPaidConsumerService = orderConsumer.NewService(
			d.OrderPaidConsumer(),
			d.OrderPaidDecoder(),
			d.OrderAssembledProducerService(),
		)
	}

	return d.orderPaidConsumerService
}

func (d *diContainer) OrderPaidConsumer() wrappedKafka.Consumer {
	if d.orderPaidConsumer == nil {
		d.orderPaidConsumer = wrappedKafkaConsumer.NewConsumer(
			d.ConsumerGroup(),
			[]string{
				config.AppConfig().OrderPaidConsumer.Topic(),
			},
			logger.Logger(),
			kafkaMiddleware.Logging(logger.Logger()),
		)
	}

	return d.orderPaidConsumer
}

func (d *diContainer) ConsumerGroup() sarama.ConsumerGroup {
	if d.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderPaidConsumer.GroupID(),
			config.AppConfig().OrderPaidConsumer.Config(),
		)
		if err != nil {
			panic(fmt.Sprintf("failed to create consumer group: %s\n", err.Error()))
		}
		closer.AddNamed("Kafka consumer group", func(ctx context.Context) error {
			return d.consumerGroup.Close()
		})

		d.consumerGroup = consumerGroup
	}

	return d.consumerGroup
}

func (d *diContainer) OrderPaidDecoder() kafkaConverter.OrderPaidDecoder {
	if d.orderPaidDecoder == nil {
		d.orderPaidDecoder = decoder.NewOrderPaidDecoder()
	}

	return d.orderPaidDecoder
}

func (d *diContainer) OrderAssembledProducerService() service.OrderProducerService {
	if d.orderAssembledProducerService == nil {
		d.orderAssembledProducerService = orderProducer.NewService(d.OrderAssembledProducer())
	}

	return d.orderAssembledProducerService
}

func (d *diContainer) OrderAssembledProducer() wrappedKafka.Producer {
	if d.orderAssembledProducer == nil {
		d.orderAssembledProducer = wrappedKafkaProducer.NewProducer(
			d.SyncProducer(),
			config.AppConfig().OrderAssembledProducer.Topic(),
			logger.Logger(),
		)
	}

	return d.orderAssembledProducer
}

func (d *diContainer) SyncProducer() sarama.SyncProducer {
	if d.syncProducer == nil {
		p, err := sarama.NewSyncProducer(
			config.AppConfig().Kafka.Brokers(),
			config.AppConfig().OrderAssembledProducer.Config(),
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

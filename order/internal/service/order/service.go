package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kirillmc/starShipsCompany/order/internal/client/grpc"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/pg"
	def "github.com/kirillmc/starShipsCompany/order/internal/service"
)

var _ def.Service = (*service)(nil)

type service struct {
	orderRepo       pg.OrderRepository
	paymentClient   grpc.PaymentClient
	inventoryClient grpc.InventoryClient

	orderPaidProducer def.OrderProducerService

	pool *pgxpool.Pool
}

func NewService(
	pool *pgxpool.Pool,
	inventoryClient grpc.InventoryClient,
	paymentClient grpc.PaymentClient,
	orderPaidProducer def.OrderProducerService,
	orderRepo pg.OrderRepository,
) *service {
	return &service{
		inventoryClient:   inventoryClient,
		paymentClient:     paymentClient,
		orderRepo:         orderRepo,
		orderPaidProducer: orderPaidProducer,
		pool:              pool,
	}
}

package order

import (
	"github.com/kirillmc/starShipsCompany/order/internal/client/grpc"
	"github.com/kirillmc/starShipsCompany/order/internal/repository"
	def "github.com/kirillmc/starShipsCompany/order/internal/service"
)

var _ def.Service = (*service)(nil)

type service struct {
	repo            repository.Repository
	paymentClient   grpc.PaymentClient
	inventoryClient grpc.InventoryClient
}

func NewService(
	paymentClient grpc.PaymentClient,
	inventoryClient grpc.InventoryClient,
	repo repository.Repository,
) *service {
	return &service{
		paymentClient: paymentClient,
		repo:          repo,
	}
}

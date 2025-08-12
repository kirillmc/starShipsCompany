package grpc

import (
	"context"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, payOrderInfo model.PayOrderInfo) (model.UUID, error)
}

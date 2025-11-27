package service

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

type Service interface {
	Get(ctx context.Context, params model.GetOrderParams) (model.Order, error)
	Update(ctx context.Context, params model.UpdateOrderParams) error
	Pay(ctx context.Context, params model.PayOrderParams) (model.TransactionUUID, error)
	Create(ctx context.Context, userUUID model.UserUUID, partsUUIDs []model.PartUUID) (model.OrderInfo, error)
	Cancel(ctx context.Context, params model.CancelOrderParams) error
}

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type OrderProducerService interface {
	ProduceOrderPaid(ctx context.Context, event model.OrderPaidEvent) error
}

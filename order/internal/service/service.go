package service

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

type Service interface {
	Get(ctx context.Context, params model.GetOrderParams) (model.Order, error)
	Pay(ctx context.Context, params model.PayOrderParams) (model.TransactionUUID, error)
	Create(ctx context.Context, userUUID model.UserUUID, partsUUIDs []model.PartUUID) (model.OrderInfo, error)
	CancelOrder(ctx context.Context, params model.CancelOrderParams) error
}

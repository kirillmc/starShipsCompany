package repository

import (
	"context"

	model "github.com/kirillmc/starShipsCompany/order/internal/model"
)

type Repository interface {
	UpdateOrder(ctx context.Context, updateOrderParams model.UpdateOrderParams) error
	Get(ctx context.Context, orderUUID model.OrderUUID) (model.Order, error)
	Create(ctx context.Context, order model.CreateOrder) (model.OrderInfo, error)
}

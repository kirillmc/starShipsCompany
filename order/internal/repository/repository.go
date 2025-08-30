package repository

import (
	"context"
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

type Repository interface {
	SetStatus(ctx context.Context, orderUUID model.OrderUUID, transactionUUID model.TransactionUUID,
		status model.OrderStatus) error
	Get(ctx context.Context, params model.GetOrderParams) (serviceModel.Order, error)
	Create(ctx context.Context, order model.Order) (serviceModel.OrderInfo, error)
}

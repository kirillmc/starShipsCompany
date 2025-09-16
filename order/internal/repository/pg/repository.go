package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
)

type OrderRepository interface {
	IndexOrderParts(ctx context.Context, orderID model.OrderID) ([]model.PartUUID, error)
	Get(ctx context.Context, orderUUID model.OrderUUID) (model.Order, error)
	Create(ctx context.Context, tx pgx.Tx, order model.CreateOrder) (model.OrderInfo, error)
	CreateOrderParts(ctx context.Context, tx pgx.Tx, orderID model.OrderID, partUUIDs []model.PartUUID) error
	UpdateOrder(ctx context.Context, updateOrderParams model.UpdateOrderParams) error
}

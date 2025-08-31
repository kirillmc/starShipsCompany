package order

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func (r *repository) SetStatus(_ context.Context, orderUUID model.OrderUUID, transactionUUID model.TransactionUUID,
	status model.OrderStatus,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[orderUUID].Status = status
	r.orders[orderUUID].TransactionUUID = transactionUUID

	return nil
}

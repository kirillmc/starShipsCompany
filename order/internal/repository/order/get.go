package order

import (
	"context"
	"fmt"

	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func (r *repository) Get(_ context.Context, orderUUID model.OrderUUID) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[orderUUID]
	if !ok {
		return model.Order{}, fmt.Errorf("failed to find order with UUID %s: %w",
			orderUUID, serviceErrors.ErrNotFound)
	}

	return converter.ToModelOrder(order), nil
}

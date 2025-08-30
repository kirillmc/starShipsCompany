package order

import (
	"context"
	"fmt"

	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func (r *repository) Get(_ context.Context, params model.GetOrderParams) (serviceModel.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[params.OrderUUID]
	if !ok {
		return serviceModel.Order{}, fmt.Errorf("failed to find order with UUID %s: %w",
			params.OrderUUID, serviceErrors.ErrNotFound)
	}

	return converter.OrderToService(order), nil
}

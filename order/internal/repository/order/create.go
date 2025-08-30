package order

import (
	"context"
	"errors"
	"fmt"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func (r *repository) Create(ctx context.Context, order model.Order) (serviceModel.OrderInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.Get(ctx, model.GetOrderParams{OrderUUID: order.OrderUUID})
	if !errors.Is(err, serviceErrors.NotFoundErr) {
		return serviceModel.OrderInfo{}, err
	}
	if err == nil {
		return serviceModel.OrderInfo{},
			fmt.Errorf("order with UUID %s already exists: %w", order.OrderUUID, serviceErrors.OnConflictErr)
	}

	r.orders[order.OrderUUID] = &model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUUIDs:       order.PartUUIDs,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   order.PaymentMethod,
		Status:          order.Status,
	}

	orderInfo := converter.ToServiceOrderInfo(r.orders[order.OrderUUID].OrderUUID, r.orders[order.OrderUUID].TotalPrice)

	return orderInfo, nil
}

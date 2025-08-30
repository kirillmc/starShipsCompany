package order

import (
	"context"

	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func (r *repository) Create(_ context.Context, order model.Order) (serviceModel.OrderInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

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

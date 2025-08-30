package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func OrderToService(order *model.Order) serviceModel.Order {
	if order == nil {
		return serviceModel.Order{}
	}

	serviceOrder := serviceModel.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUUIDs:       order.PartUUIDs,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   PaymentMethodToService(order.PaymentMethod),
		Status:          OrderStatusToService(order.Status),
	}

	return serviceOrder
}

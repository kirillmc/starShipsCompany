package converter

import (
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func ToModelOrder(order *repoModel.Order) model.Order {
	if order == nil {
		return model.Order{}
	}

	serviceOrder := model.Order{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUUIDs:       order.PartUUIDs,
		TotalPrice:      order.TotalPrice,
		TransactionUUID: order.TransactionUUID,
		PaymentMethod:   ToModelPaymentMethod(order.PaymentMethod),
		Status:          ToModelOrderStatus(order.Status),
	}

	return serviceOrder
}

func ToRepoGetOrderParams(params model.GetOrderParams) repoModel.GetOrderParams {
	return repoModel.GetOrderParams{
		OrderUUID: params.OrderUUID,
	}
}

func ToRepoCreateOrder(createOrder model.CreateOrder) repoModel.Order {
	orderMapped := repoModel.Order{
		OrderUUID:       createOrder.OrderUUID,
		UserUUID:        createOrder.UserUUID,
		PartUUIDs:       createOrder.PartsUUIDS,
		TotalPrice:      createOrder.TotalPrice,
		TransactionUUID: "",
		PaymentMethod:   "",
		Status:          repoModel.OrderStatusPendingPayment,
	}

	return orderMapped
}

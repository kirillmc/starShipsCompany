package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func OrderToGetOrderRes(order model.Order) orderV1.GetOrderResponse {
	orderAPI := orderV1.GetOrderResponse{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUUIDs,
		TotalPrice:      SetTotalPriceAPI(order.TotalPrice),
		TransactionUUID: SetTransactionUUIDAPI(order.TransactionUUID),
		PaymentMethod:   PaymentMethodToAPI(order.PaymentMethod),
		Status:          OrderStatusToAPI(order.Status),
	}

	return orderAPI
}

func GetOrderParamsToRepo(params model.GetOrderParams) repoModel.GetOrderParams {
	return repoModel.GetOrderParams{
		OrderUUID: params.OrderUUID,
	}
}

func ToCreateOrderRepo(orderUUID model.OrderUUID, userUUID model.UserUUID, partsUUIDS []model.PartUUID,
	totalPrice float64) repoModel.Order {
	repoOrder := repoModel.Order{
		OrderUUID:       orderUUID,
		UserUUID:        userUUID,
		PartUUIDs:       partsUUIDS,
		TotalPrice:      totalPrice,
		TransactionUUID: "",
		PaymentMethod:   "",
		Status:          repoModel.PENDINGPAYMENT,
	}

	return repoOrder
}

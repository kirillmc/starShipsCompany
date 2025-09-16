package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func ToAPICreateOrderResponse(orderInfo model.OrderInfo) orderV1.CreateOrderResponse {
	return orderV1.CreateOrderResponse{
		OrderUUID:  orderInfo.OrderUUID,
		TotalPrice: orderInfo.TotalPrice,
	}
}

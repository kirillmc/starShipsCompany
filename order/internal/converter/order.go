package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func ToAPIGetOrderResponse(order model.Order) orderV1.GetOrderResponse {
	orderMapped := orderV1.GetOrderResponse{
		OrderUUID:       order.OrderUUID,
		UserUUID:        order.UserUUID,
		PartUuids:       order.PartUUIDs,
		TotalPrice:      ToAPITotalPrice(order.TotalPrice),
		TransactionUUID: ToAPITransactionUUID(order.TransactionUUID),
		PaymentMethod:   ToAPIPaymentMethod(order.PaymentMethod),
		Status:          ToAPIOrderStatus(order.Status),
	}

	return orderMapped
}

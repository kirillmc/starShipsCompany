package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func ToAPIOrderStatus(status model.OrderStatus) orderV1.OptOrderStatus {
	var respStatus orderV1.OrderStatus

	switch status {
	case model.OrderStatusPendingPayment:
		respStatus = orderV1.OrderStatusPENDINGPAYMENT
	case model.OrderStatusPaid:
		respStatus = orderV1.OrderStatusPAID
	case model.OrderStatusCancelled:
		respStatus = orderV1.OrderStatusCANCELLED
	default:
		respStatus = orderV1.OrderStatusUNSPECIFIED
	}

	resp := orderV1.OptOrderStatus{
		Value: respStatus,
		Set:   true,
	}

	return resp
}

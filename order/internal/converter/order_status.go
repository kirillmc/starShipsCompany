package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func OrderStatusToRepo(status model.OrderStatus) repoModel.OrderStatus {
	switch status {
	case model.PENDINGPAYMENT:
		return repoModel.PENDINGPAYMENT
	case model.PAID:
		return repoModel.PAID
	case model.CANCELLED:
		return repoModel.CANCELLED
	default:
		return repoModel.UNSPECIFIED_ORDER_STATUS
	}
}

func OrderStatusToAPI(status model.OrderStatus) orderV1.OptOrderStatus {
	var respStatus orderV1.OrderStatus

	switch status {
	case model.PENDINGPAYMENT:
		respStatus = orderV1.OrderStatusPENDINGPAYMENT
	case model.PAID:
		respStatus = orderV1.OrderStatusPAID
	case model.CANCELLED:
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

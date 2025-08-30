package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func OrderStatusToService(status model.OrderStatus) serviceModel.OrderStatus {
	switch status {
	case model.PENDINGPAYMENT:
		return serviceModel.PENDINGPAYMENT
	case model.PAID:
		return serviceModel.PAID
	case model.CANCELLED:
		return serviceModel.CANCELLED
	default:
		return serviceModel.UNSPECIFIED_ORDER_STATUS
	}
}

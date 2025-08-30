package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func ToServiceOrderInfo(orderUUID model.OrderUUID, totalPrice float64) serviceModel.OrderInfo {
	serviceOrderInfo := serviceModel.OrderInfo{
		OrderUUID:  orderUUID,
		TotalPrice: totalPrice,
	}

	return serviceOrderInfo
}

package converter

import (
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func ToModelOrderInfo(orderUUID repoModel.OrderUUID, totalPrice repoModel.Price) model.OrderInfo {
	return model.OrderInfo{
		OrderUUID:  orderUUID,
		TotalPrice: totalPrice,
	}
}

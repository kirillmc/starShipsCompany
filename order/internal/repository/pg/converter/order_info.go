package converter

import (
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
)

func ToModelOrderInfo(orderInfo repoModel.CreatedOrderInfo) model.OrderInfo {
	return model.OrderInfo{
		ID:         orderInfo.ID,
		OrderUUID:  orderInfo.OrderUUID,
		TotalPrice: orderInfo.TotalPrice,
	}
}

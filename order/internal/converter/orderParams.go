package converter

import "github.com/kirillmc/starShipsCompany/order/internal/model"

func CancelOrderParamsToGet(params model.CancelOrderParams) model.GetOrderParams {
	res := model.GetOrderParams(params)

	return res
}

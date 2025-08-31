package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func CancelOrderParamsToGet(params model.CancelOrderParams) model.GetOrderParams {
	res := model.GetOrderParams(params)

	return res
}

func ToPayOrderParams(req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) model.PayOrderParams {
	res := model.PayOrderParams{
		OrderUUID:     params.OrderUUID.String(),
		PaymentMethod: PaymentMethodToService(req.PaymentMethod),
	}

	return res
}

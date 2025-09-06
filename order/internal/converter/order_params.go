package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func ToModelPayOrderParams(req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) model.PayOrderParams {
	return model.PayOrderParams{
		OrderUUID:     params.OrderUUID.String(),
		PaymentMethod: ToModelPaymentMethod(req.PaymentMethod),
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func ModelToPayOrderRequest(req model.PayOrderParams) *paymentV1.PayOrderRequest {
	payOrderInfo := &paymentV1.PayOrderRequest{
		OrderUuid:     req.OrderUUID,
		UserUuid:      req.UserUUID,
		PaymentMethod: PaymentMethodToProto(req.PaymentMethod),
	}

	return payOrderInfo
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func ToProtoPayOrderRequest(params model.PayOrderParams) *paymentV1.PayOrderRequest {
	payOrderInfo := &paymentV1.PayOrderRequest{
		OrderUuid:     params.OrderUUID,
		UserUuid:      params.UserUUID,
		PaymentMethod: ToProtoPaymentMethod(params.PaymentMethod),
	}
	return payOrderInfo
}

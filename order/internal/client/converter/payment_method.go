package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func PaymentMethodToProto(method model.PaymentMethod) paymentV1.PAYMENTMETHOD {
	switch method {
	case model.CARD:
		return paymentV1.PAYMENTMETHOD_CARD
	case model.SBP:
		return paymentV1.PAYMENTMETHOD_SBP
	case model.CREDITCARD:
		return paymentV1.PAYMENTMETHOD_CREDIT_CARD
	case model.INVESTORMONEY:
		return paymentV1.PAYMENTMETHOD_INVESTOR_MONEY
	default:
		return paymentV1.PAYMENTMETHOD_UNSPECIFIED
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func ToProtoPaymentMethod(method model.PaymentMethod) paymentV1.PAYMENTMETHOD {
	switch method {
	case model.PaymentMethodCard:
		return paymentV1.PAYMENTMETHOD_CARD
	case model.PaymentMethodSBP:
		return paymentV1.PAYMENTMETHOD_SBP
	case model.PaymentMethodCreditCard:
		return paymentV1.PAYMENTMETHOD_CREDIT_CARD
	case model.PaymentMethodInvestorMoney:
		return paymentV1.PAYMENTMETHOD_INVESTOR_MONEY
	default:
		return paymentV1.PAYMENTMETHOD_UNSPECIFIED
	}
}

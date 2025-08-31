package converter

import (
	"github.com/kirillmc/starShipsCompany/payment/internal/model"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func PaymentMethodToModel(method paymentV1.PAYMENTMETHOD) model.PaymentMethod {
	switch method {
	case paymentV1.PAYMENTMETHOD_CARD:
		return model.CARD
	case paymentV1.PAYMENTMETHOD_SBP:
		return model.SBP
	case paymentV1.PAYMENTMETHOD_CREDIT_CARD:
		return model.CREDITCARD
	case paymentV1.PAYMENTMETHOD_INVESTOR_MONEY:
		return model.INVESTORMONEY
	default:
		return model.UNSPECIFIED
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func PaymentMethodToAPI(method model.PaymentMethod) orderV1.OptPaymentMethod {
	var respMethod orderV1.PaymentMethod

	switch method {
	case model.CARD:
		respMethod = orderV1.PaymentMethodCARD
	case model.SBP:
		respMethod = orderV1.PaymentMethodSBP
	case model.CREDITCARD:
		respMethod = orderV1.PaymentMethodCREDITCARD
	case model.INVESTORMONEY:
		respMethod = orderV1.PaymentMethodINVESTORMONEY
	default:
		respMethod = orderV1.PaymentMethodUNSPECIFIED
	}

	resp := orderV1.OptPaymentMethod{
		Value: respMethod,
		Set:   true,
	}

	return resp
}

func PaymentMethodToService(method orderV1.PaymentMethod) model.PaymentMethod {
	var respMethod model.PaymentMethod

	switch method {
	case orderV1.PaymentMethodCARD:
		respMethod = model.CARD
	case orderV1.PaymentMethodSBP:
		respMethod = model.SBP
	case orderV1.PaymentMethodCREDITCARD:
		respMethod = model.CREDITCARD
	case orderV1.PaymentMethodINVESTORMONEY:
		respMethod = model.INVESTORMONEY
	default:
		respMethod = model.UNSPECIFIED_PAYMENT_METHOD
	}

	return respMethod
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func ToAPIPaymentMethod(method model.PaymentMethod) orderV1.OptPaymentMethod {
	var respMethod orderV1.PaymentMethod

	switch method {
	case model.PaymentMethodCard:
		respMethod = orderV1.PaymentMethodCARD
	case model.PaymentMethodSBP:
		respMethod = orderV1.PaymentMethodSBP
	case model.PaymentMethodCreditCard:
		respMethod = orderV1.PaymentMethodCREDITCARD
	case model.PaymentMethodInvestorMoney:
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

func ToModelPaymentMethod(method orderV1.PaymentMethod) model.PaymentMethod {
	var respMethod model.PaymentMethod

	switch method {
	case orderV1.PaymentMethodCARD:
		respMethod = model.PaymentMethodCard
	case orderV1.PaymentMethodSBP:
		respMethod = model.PaymentMethodSBP
	case orderV1.PaymentMethodCREDITCARD:
		respMethod = model.PaymentMethodCreditCard
	case orderV1.PaymentMethodINVESTORMONEY:
		respMethod = model.PaymentMethodInvestorMoney
	default:
		respMethod = model.PaymentMethodUnspecified
	}

	return respMethod
}

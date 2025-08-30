package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func PaymentMethodToService(status model.PaymentMethod) serviceModel.PaymentMethod {
	switch status {
	case model.CARD:
		return serviceModel.CARD
	case model.SBP:
		return serviceModel.SBP
	case model.CREDITCARD:
		return serviceModel.CREDITCARD
	case model.INVESTORMONEY:
		return serviceModel.INVESTORMONEY
	default:
		return serviceModel.UNSPECIFIED_PAYMENT_METHOD
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func ToAPITotalPrice(totalPrice model.Price) orderV1.OptFloat64 {
	return orderV1.OptFloat64{
		Value: totalPrice,
		Set:   true,
	}
}

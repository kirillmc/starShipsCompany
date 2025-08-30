package converter

import orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"

func SetTotalPriceAPI(totalPrice float64) orderV1.OptFloat64 {
	return orderV1.OptFloat64{
		Value: totalPrice,
		Set:   true,
	}
}

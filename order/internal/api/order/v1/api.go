package v1

import (
	"github.com/kirillmc/starShipsCompany/order/internal/service"
)

type api struct {
	orderService service.Service
}

func NewAPI(orderService service.Service) *api {
	return &api{
		orderService: orderService,
	}
}

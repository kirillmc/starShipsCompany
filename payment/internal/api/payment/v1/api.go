package v1

import (
	"github.com/kirillmc/starShipsCompany/payment/internal/service"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

type api struct {
	paymentV1.UnimplementedPaymentServiceServer

	paymentService service.Service
}

func NewAPI(paymentService service.Service) *api {
	return &api{
		paymentService: paymentService,
	}
}

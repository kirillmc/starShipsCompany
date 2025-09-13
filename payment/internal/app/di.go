package app

import (
	"context"
	v1 "github.com/kirillmc/starShipsCompany/payment/internal/api/payment/v1"
	"github.com/kirillmc/starShipsCompany/payment/internal/service"
	"github.com/kirillmc/starShipsCompany/payment/internal/service/payment"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

type diContainer struct {
	paymentV1API paymentV1.PaymentServiceServer

	paymentService service.Service
}

func NewDIContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) PaymentV1API(ctx context.Context) paymentV1.PaymentServiceServer {
	if d.paymentV1API == nil {
		d.paymentV1API = v1.NewAPI(d.PaymentService(ctx))
	}

	return d.paymentV1API
}

func (d *diContainer) PaymentService(ctx context.Context) service.Service {
	if d.paymentService == nil {
		d.paymentService = payment.NewService()
	}

	return d.paymentService
}

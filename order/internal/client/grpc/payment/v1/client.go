package v1

import (
	def "github.com/kirillmc/starShipsCompany/order/internal/client/grpc"
	generatedPaymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

var _ def.PaymentClient = (*client)(nil)

type client struct {
	generatedClient generatedPaymentV1.PaymentServiceClient
}

func NewClient(generatedClient generatedPaymentV1.PaymentServiceClient) *client {
	return &client{
		generatedClient: generatedClient,
	}
}

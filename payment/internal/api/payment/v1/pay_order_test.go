package v1

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/payment/internal/converter"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	var (
		ctx = context.Background()

		req = &paymentV1.PayOrderRequest{
			OrderUuid:     gofakeit.UUID(),
			UserUuid:      gofakeit.UUID(),
			PaymentMethod: paymentV1.PAYMENTMETHOD_CARD,
		}
		payOrderInfo = converter.PayOrderRequestToModel(req)

		transactionUUID = gofakeit.UUID()
	)

	s.service.On("Pay", ctx, payOrderInfo).Return(transactionUUID).Once()

	resp, err := s.api.PayOrder(ctx, req)
	s.Assert().NoError(err)
	s.Assert().NotNil(resp)
	s.Assert().Equal(transactionUUID, resp.TransactionUuid)
}

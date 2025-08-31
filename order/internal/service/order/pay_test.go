package order

import (
	"context"
	"errors"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	var (
		ctx = context.Background()

		params = model.PayOrderParams{
			OrderUUID:     gofakeit.UUID(),
			UserUUID:      gofakeit.UUID(),
			PaymentMethod: model.CARD,
		}

		getOrderParams         = model.GetOrderParams{OrderUUID: params.OrderUUID}
		foundedTransactionUUID = gofakeit.UUID()
		foundedOrder           = model.Order{
			OrderUUID: params.OrderUUID,
		}
	)

	s.repository.On("Get", ctx, converter.GetOrderParamsToRepo(getOrderParams)).
		Return(foundedOrder, nil).Once()
	s.paymentClient.On("PayOrder", ctx, params).Return(foundedTransactionUUID, nil).Once()
	s.repository.On("SetStatus", ctx, params.OrderUUID, foundedTransactionUUID,
		converter.OrderStatusToRepo(model.PAID)).Return(nil).Once()

	transactionUUID, err := s.service.Pay(ctx, params)
	s.Assert().NoError(err)
	s.Assert().Equal(foundedTransactionUUID, transactionUUID)
}

func (s *ServiceSuite) TestFailedPayUnknownOrder() {
	var (
		ctx = context.Background()

		params = model.PayOrderParams{
			OrderUUID:     gofakeit.UUID(),
			UserUUID:      gofakeit.UUID(),
			PaymentMethod: model.CARD,
		}

		getOrderParams         = model.GetOrderParams{OrderUUID: params.OrderUUID}
		foundedTransactionUUID = ""
		foundedErr             = serviceErrors.ErrNotFound
		foundedOrder           = model.Order{}
	)

	s.repository.On("Get", ctx, converter.GetOrderParamsToRepo(getOrderParams)).
		Return(foundedOrder, serviceErrors.ErrNotFound).Once()

	transactionUUID, err := s.service.Pay(ctx, params)
	s.Assert().Error(err)
	s.Assert().Equal(foundedErr, err)
	s.Assert().Equal(foundedTransactionUUID, transactionUUID)
}

func (s *ServiceSuite) TestFailedPayAlreadyPayedOrder() {
	var (
		ctx = context.Background()

		params = model.PayOrderParams{
			OrderUUID:     gofakeit.UUID(),
			UserUUID:      gofakeit.UUID(),
			PaymentMethod: model.CARD,
		}

		getOrderParams         = model.GetOrderParams{OrderUUID: params.OrderUUID}
		foundedTransactionUUID = ""
		foundedErr             = serviceErrors.ErrOnConflict
		foundedOrder           = model.Order{
			OrderUUID: params.OrderUUID,
			UserUUID:  params.UserUUID,
			Status:    model.PAID,
		}
	)

	s.repository.On("Get", ctx, converter.GetOrderParamsToRepo(getOrderParams)).
		Return(foundedOrder, nil).Once()

	transactionUUID, err := s.service.Pay(ctx, params)
	s.Assert().Error(err)
	s.Assert().True(errors.Is(err, foundedErr))
	s.Assert().Equal(foundedTransactionUUID, transactionUUID)
}

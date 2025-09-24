package order

import (
	"context"
	"errors"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestPayOrderSuccess() {
	var (
		ctx = context.Background()

		params = model.PayOrderParams{
			OrderUUID:     gofakeit.UUID(),
			UserUUID:      gofakeit.UUID(),
			PaymentMethod: model.PaymentMethodCard,
		}

		getOrderParams         = model.GetOrderParams{OrderUUID: params.OrderUUID}
		foundedTransactionUUID = gofakeit.UUID()
		foundedOrder           = model.Order{
			OrderUUID: params.OrderUUID,
		}

		updateOrderParams = model.UpdateOrderParams{
			OrderUUID:       params.OrderUUID,
			TransactionUUID: &foundedTransactionUUID,
			Status:          lo.ToPtr(model.OrderStatusPaid),
			PaymentMethod:   lo.ToPtr(params.PaymentMethod),
		}
	)

	s.orderRepository.On("Get", ctx, getOrderParams.OrderUUID).
		Return(foundedOrder, nil).Once()
	s.paymentClient.On("PayOrder", ctx, params).Return(foundedTransactionUUID, nil).Once()
	s.orderRepository.On("UpdateOrder", ctx, updateOrderParams).Return(nil).Once()
	s.orderPaidProducer.On("ProduceOrderPaid", ctx, mock.Anything).Return(nil).Once()
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
			PaymentMethod: model.PaymentMethodCard,
		}

		getOrderParams         = model.GetOrderParams{OrderUUID: params.OrderUUID}
		foundedTransactionUUID = ""
		foundedErr             = serviceErrors.ErrNotFound
		foundedOrder           = model.Order{}
	)

	s.orderRepository.On("Get", ctx, getOrderParams.OrderUUID).
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
			PaymentMethod: model.PaymentMethodCard,
		}

		getOrderParams         = model.GetOrderParams{OrderUUID: params.OrderUUID}
		foundedTransactionUUID = ""
		foundedErr             = serviceErrors.ErrOnConflict
		foundedOrder           = model.Order{
			OrderUUID: params.OrderUUID,
			UserUUID:  params.UserUUID,
			Status:    model.OrderStatusPaid,
		}
	)

	s.orderRepository.On("Get", ctx, getOrderParams.OrderUUID).
		Return(foundedOrder, nil).Once()

	transactionUUID, err := s.service.Pay(ctx, params)
	s.Assert().Error(err)
	s.Assert().True(errors.Is(err, foundedErr))
	s.Assert().Equal(foundedTransactionUUID, transactionUUID)
}

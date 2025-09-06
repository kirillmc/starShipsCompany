package v1

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/service"
	"github.com/kirillmc/starShipsCompany/order/internal/service/mocks"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPayOrder(t *testing.T) {
	t.Parallel()
	type serviceMockFunc func(t *testing.T) service.Service

	ctx := context.Background()
	transactionUUID := gofakeit.UUID()
	orderUUID := uuid.New()

	tests := []struct {
		name                   string
		req                    *orderV1.PayOrderRequest
		params                 orderV1.PayOrderParams
		err                    error
		foundedTransactionUUID model.TransactionUUID
		expectedResp           orderV1.PayOrderRes
		serviceMockFunc        serviceMockFunc
	}{
		{
			name:   "Успешная оплата заказа",
			req:    &orderV1.PayOrderRequest{PaymentMethod: orderV1.PaymentMethodCARD},
			params: orderV1.PayOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.PayOrderResponse{
				TransactionUUID: transactionUUID,
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Pay", ctx, mock.Anything).Return(transactionUUID, nil).Once()

				return mockService
			},
		},
		{
			name:   "Не найден заказ по UUID",
			req:    &orderV1.PayOrderRequest{PaymentMethod: orderV1.PaymentMethodCARD},
			params: orderV1.PayOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: serviceErrors.ErrNotFound.Error(),
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Pay", ctx, mock.Anything).Return("",
					serviceErrors.ErrNotFound).Once()

				return mockService
			},
		},
		{
			name:   "Заказ с UUID уже оплачен/отменен",
			req:    &orderV1.PayOrderRequest{PaymentMethod: orderV1.PaymentMethodCARD},
			params: orderV1.PayOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.ConflictError{
				Code:    http.StatusConflict,
				Message: serviceErrors.ErrOnConflict.Error(),
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Pay", ctx, mock.Anything).Return("",
					serviceErrors.ErrOnConflict).Once()

				return mockService
			},
		},
		{
			name:   "Внутренняя ошибка сервера",
			req:    &orderV1.PayOrderRequest{PaymentMethod: orderV1.PaymentMethodCARD},
			params: orderV1.PayOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: serviceErrors.ErrInternalServer.Error(),
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Pay", ctx, mock.Anything).Return("",
					serviceErrors.ErrInternalServer).Once()

				return mockService
			},
		},
		{
			name:   "Не передан запрос",
			req:    nil,
			params: orderV1.PayOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.UnprocessableEntityError{
				Code:    http.StatusUnprocessableEntity,
				Message: serviceErrors.ErrUnprocessableEntity.Error(),
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)

				return mockService
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			api := NewAPI(test.serviceMockFunc(t))
			res, err := api.PayOrder(ctx, test.req, test.params)
			assert.True(t, errors.Is(err, test.err))
			assert.Equal(t, test.expectedResp, res)
		})
	}
}

package v1

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/service"
	"github.com/kirillmc/starShipsCompany/order/internal/service/mocks"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	t.Parallel()
	type serviceMockFunc func(t *testing.T) service.Service

	ctx := context.Background()
	orderUUID := gofakeit.UUID()
	userUUID := gofakeit.UUID()
	var totalPrice float64 = 101

	tests := []struct {
		name                   string
		req                    *orderV1.CreateOrderRequest
		err                    error
		foundedTransactionUUID model.TransactionUUID
		expectedResp           orderV1.CreateOrderRes
		serviceMockFunc        serviceMockFunc
	}{
		{
			name: "Успешное создание заказа",
			req: &orderV1.CreateOrderRequest{
				UserUUID:  userUUID,
				PartUuids: []model.PartUUID{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
			},
			expectedResp: &orderV1.CreateOrderResponse{
				OrderUUID:  orderUUID,
				TotalPrice: totalPrice,
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Create", ctx, mock.Anything, mock.Anything).Return(model.OrderInfo{
					OrderUUID:  orderUUID,
					TotalPrice: totalPrice,
				}, nil).Once()

				return mockService
			},
		},
		{
			name: "Заказ уже существует",
			req: &orderV1.CreateOrderRequest{
				UserUUID:  userUUID,
				PartUuids: []model.PartUUID{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
			},
			expectedResp: &orderV1.ConflictError{
				Code:    http.StatusConflict,
				Message: serviceErrors.ErrOnConflict.Error(),
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Create", ctx, mock.Anything, mock.Anything).Return(model.OrderInfo{},
					serviceErrors.ErrOnConflict).Once()

				return mockService
			},
		},
		{
			name: "Внутренняя ошибка сервера",
			req: &orderV1.CreateOrderRequest{
				UserUUID:  userUUID,
				PartUuids: []model.PartUUID{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()},
			},
			expectedResp: &orderV1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: serviceErrors.ErrInternalServer.Error(),
			},
			err: nil,
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Create", ctx, mock.Anything, mock.Anything).Return(model.OrderInfo{},
					serviceErrors.ErrInternalServer).Once()

				return mockService
			},
		},
		{
			name: "Не передан запрос",
			req:  nil,
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
			res, err := api.CreateOrder(ctx, test.req)
			assert.True(t, errors.Is(err, test.err))
			assert.Equal(t, test.expectedResp, res)
		})
	}
}

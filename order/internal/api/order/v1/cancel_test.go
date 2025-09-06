package v1

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/service"
	"github.com/kirillmc/starShipsCompany/order/internal/service/mocks"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCancelOrder(t *testing.T) {
	t.Parallel()
	type serviceMockFunc func(t *testing.T) service.Service

	ctx := context.Background()
	orderUUID := uuid.New()

	tests := []struct {
		name                   string
		params                 orderV1.CancelOrderParams
		err                    error
		foundedTransactionUUID model.TransactionUUID
		expectedResp           orderV1.CancelOrderRes
		serviceMockFunc        serviceMockFunc
	}{
		{
			name:         "Заказ успешно закрыт",
			params:       orderV1.CancelOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.CancelOrderNoContent{},
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Cancel", ctx, mock.Anything).Return(nil).Once()

				return mockService
			},
		},
		{
			name:   "Не найден заказ по UUID",
			params: orderV1.CancelOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: serviceErrors.ErrNotFound.Error(),
			},
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Cancel", ctx, mock.Anything).Return(
					serviceErrors.ErrNotFound).Once()

				return mockService
			},
		},
		{
			name:   "Заказ с UUID уже оплачен/отменен",
			params: orderV1.CancelOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.ConflictError{
				Code:    http.StatusConflict,
				Message: serviceErrors.ErrOnConflict.Error(),
			},
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Cancel", ctx, mock.Anything).Return(
					serviceErrors.ErrOnConflict).Once()

				return mockService
			},
		},
		{
			name:   "Внутренняя ошибка сервера",
			params: orderV1.CancelOrderParams{OrderUUID: orderUUID},
			expectedResp: &orderV1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: serviceErrors.ErrInternalServer.Error(),
			},
			serviceMockFunc: func(t *testing.T) service.Service {
				mockService := mocks.NewService(t)
				mockService.On("Cancel", ctx, mock.Anything).Return(
					serviceErrors.ErrInternalServer).Once()

				return mockService
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			api := NewAPI(test.serviceMockFunc(t))
			res, err := api.CancelOrder(ctx, test.params)
			assert.True(t, errors.Is(err, test.err))
			assert.Equal(t, test.expectedResp, res)
		})
	}
}

package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"go.uber.org/zap"
)

func (a *api) PayOrder(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
	if req == nil {
		logger.Error(ctx, "empty request to pay order")

		return &orderV1.UnprocessableEntityError{
			Code:    http.StatusUnprocessableEntity,
			Message: serviceErrors.ErrUnprocessableEntity.Error(),
		}, nil
	}

	transactionUUID, err := a.orderService.Pay(ctx, converter.ToModelPayOrderParams(req, params))
	if err != nil {
		logger.Error(ctx, "failed to pay order", zap.Error(err))

		if errors.Is(err, serviceErrors.ErrNotFound) {
			return &orderV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: serviceErrors.ErrNotFound.Error(),
			}, nil
		}

		if errors.Is(err, serviceErrors.ErrOnConflict) {
			return &orderV1.ConflictError{
				Code:    http.StatusConflict,
				Message: serviceErrors.ErrOnConflict.Error(),
			}, nil
		}

		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: serviceErrors.ErrInternalServer.Error(),
		}, nil
	}

	return &orderV1.PayOrderResponse{TransactionUUID: transactionUUID}, nil
}

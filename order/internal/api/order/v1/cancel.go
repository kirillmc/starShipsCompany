package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"go.uber.org/zap"
)

func (a *api) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	err := a.orderService.Cancel(ctx, model.CancelOrderParams{OrderUUID: params.OrderUUID.String()})
	if err != nil {
		logger.Error(ctx, "failed to cancel order", zap.Error(err))

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

	return &orderV1.CancelOrderNoContent{}, nil
}

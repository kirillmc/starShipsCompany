package v1

import (
	"context"
	"errors"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"net/http"
)

func (a *api) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	err := a.orderService.CancelOrder(ctx, model.CancelOrderParams{OrderUUID: params.OrderUUID.String()})
	if err != nil {
		if errors.Is(err, serviceErrors.NotFoundErr) {
			return &orderV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: serviceErrors.NotFoundErr.Error(),
			}, nil
		}

		if errors.Is(err, serviceErrors.OnConflictErr) {
			return &orderV1.ConflictError{
				Code:    http.StatusConflict,
				Message: serviceErrors.OnConflictErr.Error(),
			}, nil
		}

		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: serviceErrors.InternalServerErr.Error(),
		}, nil
	}

	return &orderV1.CancelOrderNoContent{}, nil
}

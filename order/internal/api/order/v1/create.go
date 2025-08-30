package v1

import (
	"context"
	"errors"
	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"net/http"
)

func (a *api) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	if req == nil {
		return &orderV1.UnprocessableEntityError{
			Code:    http.StatusUnprocessableEntity,
			Message: serviceErrors.UnprocessableEntityErr.Error(),
		}, nil
	}

	orderInfo, err := a.orderService.Create(ctx, req.UserUUID, req.PartUuids)
	if err != nil {
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

	resp := converter.OrderInfoToCreateOrderResponse(orderInfo)

	return &resp, nil
}

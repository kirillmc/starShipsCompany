package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func (a *api) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	if req == nil {
		return &orderV1.UnprocessableEntityError{
			Code:    http.StatusUnprocessableEntity,
			Message: serviceErrors.ErrUnprocessableEntity.Error(),
		}, nil
	}

	orderInfo, err := a.orderService.Create(ctx, req.UserUUID, req.PartUuids)
	if err != nil {
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

	resp := converter.ToAPICreateOrderResponse(orderInfo)
	return &resp, nil
}

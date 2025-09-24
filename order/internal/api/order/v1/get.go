package v1

import (
	"context"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"go.uber.org/zap"
)

func (a *api) GetOrder(ctx context.Context, params orderV1.GetOrderParams) (orderV1.GetOrderRes, error) {
	order, err := a.orderService.Get(ctx, model.GetOrderParams{OrderUUID: params.OrderUUID.String()})
	if err != nil {
		logger.Error(ctx, "failed to get order", zap.Error(err))

		if errors.Is(err, serviceErrors.ErrNotFound) {
			return &orderV1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: serviceErrors.ErrNotFound.Error(),
			}, nil
		}

		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: serviceErrors.ErrInternalServer.Error(),
		}, nil
	}

	orderResp := converter.ToAPIGetOrderResponse(order)
	return &orderResp, nil
}

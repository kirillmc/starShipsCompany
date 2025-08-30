package v1

import (
	"context"
	"errors"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
	"net/http"
)

func (a *api) PayOrder(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
	if req == nil {
		return &orderV1.UnprocessableEntityError{
			Code:    http.StatusUnprocessableEntity,
			Message: serviceErrors.UnprocessableEntityErr.Error(),
		}, nil
	}

	transactionUUID, err := a.orderService.Pay(ctx, model.PayOrderParams{OrderUUID: params.OrderUUID.String()})
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

	return &orderV1.PayOrderResponse{TransactionUUID: transactionUUID}, nil
}

package v1

import (
	"context"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func (a *api) CancelOrder(ctx context.Context, params orderV1.CancelOrderParams) (orderV1.CancelOrderRes, error) {
	order, err := a.storage.getOrder(params.OrderUUID.String())
	if err != nil {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	if order.Status.Value == orderV1.OrderStatusPAID {
		return &orderV1.ConflictError{
			Code:    409,
			Message: "Заказ уже отменен",
		}, nil
	}

	a.storage.setOrderStatus(params.OrderUUID.String(), orderV1.OrderStatusCANCELLED)

	return &orderV1.CancelOrderNoContent{}, nil
}

package v1

import (
	"context"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func (a *api) GetOrder(ctx context.Context, params orderV1.GetOrderParams) (orderV1.GetOrderRes, error) {
	order, err := a.storage.getOrder(params.OrderUUID.String())
	if err != nil {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	return order, nil
}

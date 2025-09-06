package order

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *service) Get(ctx context.Context, params model.GetOrderParams) (model.Order, error) {
	order, err := s.repo.Get(ctx, params.OrderUUID)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

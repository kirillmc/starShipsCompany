package order

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *service) Get(ctx context.Context, params model.GetOrderParams) (model.Order, error) {
	order, err := s.orderRepo.Get(ctx, params.OrderUUID)
	if err != nil {
		return model.Order{}, err
	}

	partUUIDs, err := s.orderPartRepo.Index(ctx, order.ID)
	if err != nil {
		return model.Order{}, err
	}
	order.PartUUIDs = partUUIDs

	return order, nil
}

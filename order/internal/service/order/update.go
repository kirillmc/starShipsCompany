package order

import (
	"context"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *service) Update(ctx context.Context, params model.UpdateOrderParams) error {
	err := s.orderRepo.UpdateOrder(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

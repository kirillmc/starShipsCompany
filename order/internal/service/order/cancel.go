package order

import (
	"context"
	"fmt"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/samber/lo"
)

func (s *service) Cancel(ctx context.Context, params model.CancelOrderParams) error {
	order, err := s.orderRepo.Get(ctx, params.OrderUUID)
	if err != nil {
		return err
	}

	if order.Status == model.OrderStatusPaid {
		return fmt.Errorf("order is aleready paid: %w", serviceErrors.ErrOnConflict)
	}

	if order.Status == model.OrderStatusCancelled {
		return fmt.Errorf("order is aleready cancelled: %w", serviceErrors.ErrOnConflict)
	}

	updateOrderParams := model.UpdateOrderParams{
		OrderUUID: params.OrderUUID,
		Status:    lo.ToPtr(model.OrderStatusCancelled),
	}
	err = s.orderRepo.UpdateOrder(ctx, updateOrderParams)
	if err != nil {
		return err
	}

	return nil
}

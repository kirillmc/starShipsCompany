package order

import (
	"context"
	"fmt"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (s *service) Cancel(ctx context.Context, params model.CancelOrderParams) error {
	order, err := s.orderRepo.Get(ctx, params.OrderUUID)
	if err != nil {
		return err
	}

	if order.Status == model.OrderStatusPaid {
		logger.Error(ctx, fmt.Sprintf("order with UUID %s is already paid", params.OrderUUID), zap.Error(err))
		return fmt.Errorf("order is already paid: %w", serviceErrors.ErrOnConflict)
	}

	if order.Status == model.OrderStatusCancelled {
		logger.Error(ctx, fmt.Sprintf("order with UUID %s is already cancelled",
			params.OrderUUID), zap.Error(err))
		return fmt.Errorf("order is already cancelled: %w", serviceErrors.ErrOnConflict)
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

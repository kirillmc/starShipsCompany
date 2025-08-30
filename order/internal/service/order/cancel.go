package order

import (
	"context"
	"fmt"

	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *service) CancelOrder(ctx context.Context, params model.CancelOrderParams) error {
	order, err := s.Get(ctx, converter.CancelOrderParamsToGet(params))
	if err != nil {
		return err
	}

	if order.Status == model.PAID {
		return fmt.Errorf("order is aleready paid: %w", serviceErrors.ErrOnConflict)
	}

	if order.Status == model.CANCELLED {
		return fmt.Errorf("order is aleready cancelled: %w", serviceErrors.ErrOnConflict)
	}

	err = s.repo.SetStatus(ctx, params.OrderUUID, order.TransactionUUID, converter.OrderStatusToRepo(model.CANCELLED))
	if err != nil {
		return err
	}

	return nil
}

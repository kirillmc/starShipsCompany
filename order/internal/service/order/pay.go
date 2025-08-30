package order

import (
	"context"
	"fmt"

	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *service) Pay(ctx context.Context, params model.PayOrderParams) (model.TransactionUUID, error) {
	order, err := s.Get(ctx, model.GetOrderParams{OrderUUID: params.OrderUUID})
	if err != nil {
		return "", err
	}

	if order.Status == model.PAID {
		return "", fmt.Errorf("order is aleready paid: %w", serviceErrors.ErrOnConflict)
	}

	if order.Status == model.CANCELLED {
		return "", fmt.Errorf("order is aleready cancelled: %w", serviceErrors.ErrOnConflict)
	}

	transactionUUID, err := s.paymentClient.PayOrder(ctx, params)
	if err != nil {
		return "", fmt.Errorf("failed to pay order: %w", err)
	}

	err = s.repo.SetStatus(ctx, params.OrderUUID, transactionUUID, converter.OrderStatusToRepo(model.PAID))
	if err != nil {
		return "", err
	}

	return transactionUUID, nil
}

package order

import (
	"context"
	"fmt"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/samber/lo"
)

func (s *service) Pay(ctx context.Context, params model.PayOrderParams) (model.TransactionUUID, error) {
	order, err := s.repo.Get(ctx, params.OrderUUID)
	if err != nil {
		return "", err
	}

	if order.Status == model.OrderStatusPaid {
		return "", fmt.Errorf("order is aleready paid: %w", serviceErrors.ErrOnConflict)
	}

	if order.Status == model.OrderStatusCancelled {
		return "", fmt.Errorf("order is aleready cancelled: %w", serviceErrors.ErrOnConflict)
	}

	transactionUUID, err := s.paymentClient.PayOrder(ctx, params)
	if err != nil {
		return "", fmt.Errorf("failed to pay order: %w", err)
	}

	updateOrderParams := model.UpdateOrderParams{
		OrderUUID:       params.OrderUUID,
		TransactionUUID: &transactionUUID,
		Status:          lo.ToPtr(model.OrderStatusPaid),
	}
	err = s.repo.UpdateOrder(ctx, updateOrderParams)
	if err != nil {
		return "", err
	}

	return transactionUUID, nil
}

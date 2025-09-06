package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func (s *service) Create(ctx context.Context, userUUID model.UserUUID,
	partsUUIDs []model.PartUUID,
) (model.OrderInfo, error) {
	parts, err := s.inventoryClient.ListParts(ctx, model.PartsFilter{UUIDs: partsUUIDs})
	if err != nil {
		return model.OrderInfo{}, err
	}

	if len(parts) < len(partsUUIDs) {
		return model.OrderInfo{}, fmt.Errorf("not enough parts: %w", serviceErrors.ErrInternalServer)
	}

	var totalPrice model.Price
	partsUUIDS := make([]model.PartUUID, 0, len(parts))
	for _, part := range parts {
		totalPrice += part.Price
		partsUUIDS = append(partsUUIDS, part.UUID)
	}

	orderUUID := uuid.NewString()

	_, err = s.Get(ctx, model.GetOrderParams{OrderUUID: orderUUID})
	if !errors.Is(err, serviceErrors.ErrNotFound) {
		return model.OrderInfo{}, err
	}
	if err == nil {
		return model.OrderInfo{},
			fmt.Errorf("order with UUID %s already exists: %w", orderUUID, serviceErrors.ErrOnConflict)
	}

	createOrderInfo := model.CreateOrder{
		OrderUUID:  orderUUID,
		UserUUID:   userUUID,
		PartsUUIDS: partsUUIDS,
		TotalPrice: totalPrice,
	}
	orderInfo, err := s.repo.Create(ctx, createOrderInfo)
	if err != nil {
		return model.OrderInfo{}, err
	}

	return orderInfo, nil
}

package order

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/order/internal/converter"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/error"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (s *service) Create(ctx context.Context, userUUID model.UserUUID,
	partsUUIDs []model.PartUUID) (model.OrderInfo, error) {
	parts, err := s.inventoryClient.ListParts(ctx, model.PartsFilter{UUIDs: partsUUIDs})
	if err != nil {
		return model.OrderInfo{}, err
	}

	if len(parts) < len(partsUUIDs) {
		return model.OrderInfo{}, fmt.Errorf("not enough parts: %w", serviceErrors.InternalServerErr)
	}

	var totalPrice float64
	partsUUIDS := make([]string, 0, len(parts))
	for _, part := range parts {
		totalPrice += part.Price
		partsUUIDS = append(partsUUIDS, part.UUID)
	}

	orderUUID := uuid.NewString()

	orderInfo, err := s.repo.Create(ctx, converter.ToCreateOrderRepo(orderUUID, userUUID, partsUUIDS, totalPrice))
	if err != nil {
		return model.OrderInfo{}, err
	}

	return orderInfo, nil
}

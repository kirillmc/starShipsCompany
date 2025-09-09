package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"

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

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return model.OrderInfo{}, fmt.Errorf("%w: ошибка начала транзакции: %s",
			serviceErrors.ErrInternalServer, err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	createOrderInfo := model.CreateOrder{
		OrderUUID:  orderUUID,
		UserUUID:   userUUID,
		TotalPrice: totalPrice,
	}
	orderInfo, err := s.orderRepo.Create(ctx, tx, createOrderInfo)
	if err != nil {
		return model.OrderInfo{}, err
	}

	err = s.orderPartRepo.Create(ctx, tx, orderInfo.ID, partsUUIDS)
	if err != nil {
		return model.OrderInfo{}, err
	}

	return orderInfo, nil
}

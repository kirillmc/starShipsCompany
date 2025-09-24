package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
		return model.OrderInfo{}, fmt.Errorf("%w: failed to start tx: %s",
			serviceErrors.ErrInternalServer, err.Error())
	}
	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			err = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
			if err == nil {
				return
			}
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

	err = s.orderRepo.CreateOrderParts(ctx, tx, orderInfo.ID, partsUUIDS)
	if err != nil {
		return model.OrderInfo{}, err
	}

	return orderInfo, nil
}

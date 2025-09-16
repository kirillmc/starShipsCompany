package order

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

func (r *repository) CreateOrderParts(ctx context.Context, tx pgx.Tx, orderID model.OrderID,
	partUUIDs []model.PartUUID,
) error {
	const op = "CreateOrderParts"

	insertBuilder := sq.Insert(ordersPartsTable).PlaceholderFormat(sq.Dollar).Columns(orderIDColumn, partUUIDColumn)

	for _, partUUID := range partUUIDs {
		insertBuilder = insertBuilder.Values(orderID, partUUID)
	}

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to build %s query", op), zap.Error(err))
		return fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to execute %s query", op), zap.Error(err))
		return fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	return nil
}

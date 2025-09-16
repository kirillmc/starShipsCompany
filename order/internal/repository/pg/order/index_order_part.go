package order

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

func (r *repository) IndexOrderParts(ctx context.Context, orderID model.OrderID) ([]model.PartUUID, error) {
	const op = "IndexOrderParts"

	selectBuilder := sq.Select(partUUIDColumn).
		From(ordersPartsTable).
		Where(sq.Eq{orderIDColumn: orderID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to build %s query", op), zap.Error(err))
		return []model.PartUUID{}, fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	var partUUIDs []repoModel.PartUUID
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to execute %s query", op), zap.Error(err))
		return []model.PartUUID{}, fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	for rows.Next() {
		var partUUID repoModel.PartUUID
		err = rows.Scan(&partUUID)
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("failed during scanning values after %s query", op), zap.Error(err))
			return []model.PartUUID{}, fmt.Errorf("%w: failed during scanning values after %s query: %s",
				serviceErrors.ErrInternalServer, op, err.Error())
		}

		partUUIDs = append(partUUIDs, partUUID)
	}

	return partUUIDs, nil
}

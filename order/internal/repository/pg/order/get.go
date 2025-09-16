package order

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/pg/converter"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

func (r *repository) Get(ctx context.Context, orderUUID model.OrderUUID) (model.Order, error) {
	const op = "Get"

	selectBuilder := sq.Select("o.id", "o.order_uuid", "o.user_uuid", "o.transaction_uuid",
		"o.total_price", "o.payment_method", "o.status", "op.part_uuid").
		From("orders o").
		LeftJoin("orders_parts op on o.id = op.order_id").
		Where(sq.Eq{"o.order_uuid": orderUUID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to build %s query", op), zap.Error(err))
		return model.Order{}, fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to execute %s query", op), zap.Error(err))
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Order{}, fmt.Errorf("%w: failed to execute %s query: %s",
				serviceErrors.ErrNotFound, op, err.Error())
		}

		return model.Order{}, fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	res, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[repoModel.OrderWthPart])
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed during scanning values after %s query", op), zap.Error(err))
		return model.Order{}, fmt.Errorf("%w: failed during scanning values after %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}
	if len(res) == 0 {
		logger.Error(ctx, fmt.Sprintf("failed during scanning values after %s query", op), zap.Error(err))
		return model.Order{}, fmt.Errorf("%w: failed during scanning values after %s query",
			serviceErrors.ErrNotFound, op)
	}

	orderService, err := converter.ToModelOrder(res)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed during converting to service model after %s query", op),
			zap.Error(err))
		return model.Order{}, err
	}

	return orderService, nil
}

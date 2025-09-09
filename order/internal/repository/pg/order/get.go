package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/pg/converter"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"

	sq "github.com/Masterminds/squirrel"
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	serviceErrors "github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func (r *repository) Get(ctx context.Context, orderUUID model.OrderUUID) (model.Order, error) {
	const op = "Get"

	selectBuilder := sq.Select(idColumn, orderUUIDColumn, userUUIDColumn, transactionUUIDColumn,
		totalPriceColumn, paymentMethodColumn, statusColumn).
		From(tableName).
		Where(sq.Eq{orderUUID: orderUUID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return model.Order{}, fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err)
	}

	var order repoModel.Order
	err = r.pool.QueryRow(ctx, query, args...).
		Scan(&order.ID, &order.OrderUUID, &order.UserUUID, &order.TransactionUUID, &order.TotalPrice, &order.PaymentMethod, &order.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Order{}, fmt.Errorf("%w: failed to execute %s query: %s",
				serviceErrors.ErrNotFound, op, err)
		}

		return model.Order{}, fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err)
	}

	return converter.ToModelOrder(&order), nil
}

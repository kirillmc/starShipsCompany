package orderPart

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func (r *repository) Create(ctx context.Context, tx pgx.Tx, orderID model.OrderID, partUUIDs []model.PartUUID) error {
	const op = "Create"

	insertBuilder := sq.Insert(tableName).PlaceholderFormat(sq.Dollar).Columns(orderIDColumn, partUUIDColumn)

	for _, partUUID := range partUUIDs {
		insertBuilder = insertBuilder.Values(orderID, partUUID)
	}

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err)
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err)
	}

	return nil
}

package orderPart

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func (r *repository) Index(ctx context.Context, orderID model.OrderID) ([]model.PartUUID, error) {
	const op = "Index"

	selectBuilder := sq.Select(partUUIDColumn).
		From(tableName).
		Where(sq.Eq{orderIDColumn: orderID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return []model.PartUUID{}, fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err)
	}

	var partUUIDs []repoModel.PartUUID
	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return []model.PartUUID{}, fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err)
	}

	for rows.Next() {
		var partUUID repoModel.PartUUID
		err = rows.Scan(&partUUID)
		if err != nil {
			return []model.PartUUID{}, fmt.Errorf("%w: failed during scanning values after %s query: %s",
				serviceErrors.ErrInternalServer, op, err)
		}

		partUUIDs = append(partUUIDs, partUUID)
	}

	return partUUIDs, nil
}

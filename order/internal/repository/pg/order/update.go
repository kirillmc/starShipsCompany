package order

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/pg/converter"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func (r *repository) UpdateOrder(ctx context.Context, updateOrderParams model.UpdateOrderParams) error {
	const op = "UpdateOrder"

	updateOrderParamsRepo := converter.ToRepoUpdateOrderParams(updateOrderParams)

	updateBuilder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{orderUUIDColumn: updateOrderParamsRepo.OrderUUID})
	updateBuilder = applyUpdateFilter(updateBuilder, updateOrderParamsRepo)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	return nil
}

func applyUpdateFilter(updateBuilder sq.UpdateBuilder, updateParams repoModel.UpdateOrderParams) sq.UpdateBuilder {
	if updateParams.UserUUID != nil {
		updateBuilder = updateBuilder.Set(userUUIDColumn, updateParams.UserUUID)
	}

	if updateParams.TotalPrice != nil {
		updateBuilder = updateBuilder.Set(totalPriceColumn, updateParams.TotalPrice)
	}
	if updateParams.TransactionUUID != nil {
		updateBuilder = updateBuilder.Set(transactionUUIDColumn, updateParams.TransactionUUID)
	}
	if updateParams.PaymentMethod != nil {
		updateBuilder = updateBuilder.Set(paymentMethodColumn, updateParams.PaymentMethod)
	}
	if updateParams.Status != nil {
		updateBuilder = updateBuilder.Set(statusColumn, updateParams.Status)
	}

	return updateBuilder
}

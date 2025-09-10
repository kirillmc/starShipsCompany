package order

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	converter "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/converter"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

const returningPrefix = "RETURNING %s, %s, %s"

func (r *repository) Create(ctx context.Context, tx pgx.Tx, order model.CreateOrder) (model.OrderInfo, error) {
	const op = "Create"

	orderRepo := converter.ToRepoCreateOrder(order)
	insertBuilder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(orderUUIDColumn, userUUIDColumn, transactionUUIDColumn,
			totalPriceColumn, paymentMethodColumn, statusColumn).
		Values(orderRepo.OrderUUID, orderRepo.UserUUID, orderRepo.TransactionUUID,
			orderRepo.TotalPrice, orderRepo.PaymentMethod, orderRepo.Status).
		Suffix(fmt.Sprintf(returningPrefix, idColumn, orderUUIDColumn, totalPriceColumn))

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return model.OrderInfo{}, fmt.Errorf("%w: failed to build %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	var createdOrderInfo repoModel.CreatedOrderInfo
	err = tx.QueryRow(ctx, query, args...).
		Scan(&createdOrderInfo.ID, &createdOrderInfo.OrderUUID, &createdOrderInfo.TotalPrice)
	if err != nil {
		return model.OrderInfo{}, fmt.Errorf("%w: failed to execute %s query: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	return converter.ToModelOrderInfo(createdOrderInfo), nil
}

package orderPart

import (
	"github.com/jackc/pgx/v5/pgxpool"
	def "github.com/kirillmc/starShipsCompany/order/internal/repository/pg"
)

const (
	tableName = "orders_parts"

	partUUIDColumn = "part_uuid"
	orderIDColumn  = "order_id"
)

var _ def.OrderPartRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

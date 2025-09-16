package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
	def "github.com/kirillmc/starShipsCompany/order/internal/repository/pg"
)

const (
	idColumn              = "id"
	orderUUIDColumn       = "order_uuid"
	userUUIDColumn        = "user_uuid"
	partUUIDColumn        = "part_uuid"
	transactionUUIDColumn = "transaction_uuid"
	totalPriceColumn      = "total_price"
	paymentMethodColumn   = "payment_method"
	statusColumn          = "status"
	orderIDColumn         = "order_id"

	ordersPartsTable = "orders_parts"
	tableName        = "orders"
)

var _ def.OrderRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

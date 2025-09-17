package model

type (
	OrderID = uint64

	OrderUUID       = string
	UserUUID        = string
	PartUUID        = string
	TransactionUUID = string

	Price = float64
)

type Order struct {
	ID              OrderID          `db:"id"`
	OrderUUID       OrderUUID        `db:"order_uuid"`
	UserUUID        UserUUID         `db:"user_uuid"`
	TotalPrice      Price            `db:"total_price"`
	TransactionUUID *TransactionUUID `db:"transaction_uuid"`
	PaymentMethod   PaymentMethod    `db:"payment_method"`
	Status          OrderStatus      `db:"status"`
}

type OrderWthPart struct {
	Order
	PartUUID *PartUUID `db:"part_uuid"`
}

type CreatedOrderInfo struct {
	ID         OrderID   `json:"id"`
	OrderUUID  OrderUUID `json:"order_uuid"`
	TotalPrice Price     `json:"total_price"`
}

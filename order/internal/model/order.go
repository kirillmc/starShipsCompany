package model

type (
	OrderID = uint64

	OrderUUID       = string
	UserUUID        = string
	TransactionUUID = string

	Price = float64
)

type Order struct {
	ID              OrderID          `json:"id"`
	OrderUUID       OrderUUID        `json:"order_uuid"`
	UserUUID        UserUUID         `json:"user_uuid"`
	PartUUIDs       []PartUUID       `json:"part_uuids"`
	TotalPrice      Price            `json:"total_price"`
	TransactionUUID *TransactionUUID `json:"transaction_uuid"`
	PaymentMethod   PaymentMethod    `json:"payment_method"`
	Status          OrderStatus      `json:"status"`
}

type CreateOrder struct {
	OrderUUID  OrderUUID
	UserUUID   UserUUID
	TotalPrice Price
}

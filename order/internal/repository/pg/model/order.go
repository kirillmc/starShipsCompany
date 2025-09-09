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
	ID              OrderID         `json:"id"`
	OrderUUID       OrderUUID       `json:"order_uuid"`
	UserUUID        UserUUID        `json:"user_uuid"`
	TotalPrice      Price           `json:"total_price"`
	TransactionUUID TransactionUUID `json:"transaction_uuid"`
	PaymentMethod   PaymentMethod   `json:"payment_method"`
	Status          OrderStatus     `json:"status"`
}

type CreatedOrderInfo struct {
	ID         OrderID   `json:"id"`
	OrderUUID  OrderUUID `json:"order_uuid"`
	TotalPrice Price     `json:"total_price"`
}

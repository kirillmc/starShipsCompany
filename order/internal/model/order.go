package model

type OrderUUID = string
type UserUUID = string
type TransactionUUID = string

type Order struct {
	OrderUUID       OrderUUID       `json:"order_uuid"`
	UserUUID        UserUUID        `json:"user_uuid"`
	PartUUIDs       []PartUUID      `json:"part_uuids"`
	TotalPrice      float64         `json:"total_price"`
	TransactionUUID TransactionUUID `json:"transaction_uuid"`
	PaymentMethod   PaymentMethod   `json:"payment_method"`
	Status          OrderStatus     `json:"status"`
}

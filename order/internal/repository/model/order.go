package model

type UUID = string

type Order struct {
	OrderUUID       UUID             `json:"order_uuid"`
	UserUUID        UUID             `json:"user_uuid"`
	PartUuids       []UUID           `json:"part_uuids"`
	TotalPrice      float64          `json:"total_price"`
	TransactionUUID UUID             `json:"transaction_uuid"`
	PaymentMethod   OptPaymentMethod `json:"payment_method"`
	Status          OptOrderStatus   `json:"status"`
}

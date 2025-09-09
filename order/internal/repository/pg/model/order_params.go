package model

type GetOrderParams struct {
	OrderUUID OrderUUID
}

type CancelOrderParams struct {
	OrderUUID OrderUUID
}

type UpdateOrderParams struct {
	OrderUUID       OrderUUID        `json:"order_uuid"`
	UserUUID        *UserUUID        `json:"user_uuid"`
	TotalPrice      *Price           `json:"total_price"`
	TransactionUUID *TransactionUUID `json:"transaction_uuid"`
	PaymentMethod   *PaymentMethod   `json:"payment_method"`
	Status          *OrderStatus     `json:"status"`
}

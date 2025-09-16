package model

type GetOrderParams struct {
	OrderUUID OrderUUID
}

type CancelOrderParams struct {
	OrderUUID OrderUUID
}

type PayOrderParams struct {
	OrderUUID     OrderUUID
	PaymentMethod PaymentMethod
	UserUUID      UserUUID
}

type UpdateOrderParams struct {
	OrderUUID       OrderUUID        `json:"order_uuid"`
	UserUUID        *UserUUID        `json:"user_uuid"`
	PartUUIDs       []PartUUID       `json:"part_uuids"`
	TotalPrice      *Price           `json:"total_price"`
	TransactionUUID *TransactionUUID `json:"transaction_uuid"`
	PaymentMethod   *PaymentMethod   `json:"payment_method"`
	Status          *OrderStatus     `json:"status"`
}

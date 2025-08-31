package model

type OrderStatus string

const (
	UNSPECIFIED_ORDER_STATUS OrderStatus = "UNSPECIFIED"
	PENDINGPAYMENT           OrderStatus = "PENDING_PAYMENT"
	PAID                     OrderStatus = "PAID"
	CANCELLED                OrderStatus = "CANCELLED"
)

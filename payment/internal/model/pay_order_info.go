package model

type (
	OrderUUID       = string
	UserUUID        = string
	TransactionUUID = string
)

type PayOrderInfo struct {
	OrderUUID     OrderUUID
	UserUUID      UserUUID
	PaymentMethod PaymentMethod
}

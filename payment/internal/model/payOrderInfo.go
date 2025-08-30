package model

type OrderUUID = string
type UserUUID = string
type TransactionUUID = string

type PayOrderInfo struct {
	OrderUUID     OrderUUID
	UserUUID      UserUUID
	PaymentMethod PaymentMethod
}

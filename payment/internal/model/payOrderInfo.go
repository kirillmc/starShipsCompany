package model

type PayOrderInfo struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod PaymentMethod
}

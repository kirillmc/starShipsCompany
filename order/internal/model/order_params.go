package model

type GetOrderParams struct {
	OrderUUID OrderUUID
}

type CancelOrderParams struct {
	OrderUUID OrderUUID
}

type PayOrderParams struct {
	OrderUUID     OrderUUID
	UserUUID      UserUUID
	PaymentMethod PaymentMethod
}

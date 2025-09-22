package model

type OrderAssembledEvent struct {
	EventUUID    string
	OrderUUID    string
	UserUUID     string
	BuildTimeSec int64
}

type OrderPaidEvent struct {
	EventUUID       string
	OrderUUID       string
	UserUUID        string
	PaymentMethod   string
	TransactionUUID string
}

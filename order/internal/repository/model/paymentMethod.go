package model

type PaymentMethod string

const (
	UNSPECIFIED   PaymentMethod = "UNSPECIFIED"
	CARD          PaymentMethod = "CARD"
	SBP           PaymentMethod = "SBP"
	CREDITCARD    PaymentMethod = "CREDIT_CARD"
	INVESTORMONEY PaymentMethod = "INVESTOR_MONEY"
)

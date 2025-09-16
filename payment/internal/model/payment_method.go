package model

type PaymentMethod string

const (
	Unspecified   PaymentMethod = "UNSPECIFIED"
	Card          PaymentMethod = "CARD"
	SBP           PaymentMethod = "SBP"
	CreditCard    PaymentMethod = "CREDITCARD"
	InvestorMoney PaymentMethod = "INVESTORMONEY"
)

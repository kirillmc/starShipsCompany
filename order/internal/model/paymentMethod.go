package model

type PaymentMethod int32

const (
	UNSPECIFIED_METHOD PaymentMethod = iota
	CARD
	SBP
	CREDITCARD
	INVESTORMONEY
)

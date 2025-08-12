package model

type PaymentMethod int32

const (
	UNSPECIFIED PaymentMethod = iota
	CARD
	SBP
	CREDITCARD
	INVESTORMONEY
)

///go:build integration

package model

type Category string

const (
	Unspecified Category = "UNSPECIFIED"
	Engine      Category = "ENGINE"
	Fuel        Category = "FUEL"
	Porthole    Category = "PORTHOLE"
	Wing        Category = "WING"
)

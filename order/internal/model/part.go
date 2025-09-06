package model

import "time"

type (
	UUID     = string
	PartUUID = string
)

type Part struct {
	UUID          PartUUID
	Name          string
	Description   string
	Price         Price
	StockQuantity int64
	Category      Category
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []string
	Metadata      map[string]interface{}
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

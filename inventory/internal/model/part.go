package model

import (
	"time"
)

type (
	PartUUID = string
	Tag      = string
)

type Part struct {
	UUID          PartUUID
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []Tag
	Metadata      map[string]interface{}
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

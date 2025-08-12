package model

import "time"

type UUID = string

type Part struct {
	UUID          UUID
	Name          string
	Description   string
	Price         float64
	StockQuantity int64
	Category      Category
	Dimensions    *Dimensions
	Manufacturer  *Manufacturer
	Tags          []string
	Metadata      map[string]interface{}
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

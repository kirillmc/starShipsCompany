//go:build integration

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PartUUID = string

type Part struct {
	ID            primitive.ObjectID     `bson:"_id,omitempty"`
	UUID          PartUUID               `bson:"uuid"`
	Name          string                 `bson:"name"`
	Description   string                 `bson:"description"`
	Price         float64                `bson:"price"`
	StockQuantity int64                  `bson:"stock_quantity"`
	Category      Category               `bson:"category"`
	Dimensions    *Dimensions            `bson:"dimensions,omitempty"`
	Manufacturer  *Manufacturer          `bson:"manufacturer,omitempty"`
	Tags          []string               `bson:"tags,omitempty"`
	Metadata      map[string]interface{} `bson:"metadata,omitempty"`
	CreatedAt     *time.Time             `bson:"created_at,omitempty"`
	UpdatedAt     *time.Time             `bson:"updated_at,omitempty"`
}

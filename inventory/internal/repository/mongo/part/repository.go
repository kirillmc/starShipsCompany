package part

import (
	"context"
	def "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

const partCollectionName = "part"

var _ def.Repository = (*repository)(nil)

type repository struct {
	collection *mongo.Collection
}

func NewRepository(ctx context.Context, db *mongo.Database) (*repository, error) {
	collection := db.Collection(partCollectionName)

	err := setDefaultPartsMap(ctx, collection)
	if err != nil {
		return &repository{}, err
	}

	return &repository{
		collection: collection,
	}, nil
}

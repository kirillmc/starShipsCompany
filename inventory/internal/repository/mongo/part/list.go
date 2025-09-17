package part

import (
	"context"
	"fmt"
	"log"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/converter"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) List(ctx context.Context) ([]*model.Part, error) {
	const op = "List"
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return []*model.Part{},
			fmt.Errorf("%w: failed to execute %s: %s", serviceErrors.ErrInternalServer, op, err.Error())
	}
	defer func() {
		cerr := cursor.Close(ctx)
		if cerr != nil {
			log.Printf("failed to close cursor: %v\n", cerr.Error())
		}
	}()

	var parts []*repoModel.Part
	err = cursor.All(ctx, &parts)
	if err != nil {
		return []*model.Part{},
			fmt.Errorf("%w: failed to execute %s: %s", serviceErrors.ErrInternalServer, op, err.Error())
	}

	return converter.ToModelParts(parts), nil
}

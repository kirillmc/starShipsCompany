package part

import (
	"context"
	"fmt"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/converter"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	"go.mongodb.org/mongo-driver/bson"
)

const partUUIDField = "uuid"

func (r *repository) Get(ctx context.Context, partUUID model.PartUUID) (*model.Part, error) {
	const op = "Get"

	var part *repoModel.Part
	err := r.collection.FindOne(ctx, bson.M{partUUIDField: partUUID}).Decode(&part)
	if err != nil {
		return &model.Part{}, fmt.Errorf("%w: failed to execute %s: %s",
			serviceErrors.ErrInternalServer, op, err.Error())
	}

	return converter.ToModelPart(part), nil
}

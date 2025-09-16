package part

import (
	"context"
	"fmt"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/converter"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

const partUUIDField = "uuid"

func (r *repository) Get(ctx context.Context, partUUID model.PartUUID) (*model.Part, error) {
	const op = "Get"

	var part *repoModel.Part
	err := r.collection.FindOne(ctx, bson.M{partUUIDField: partUUID}).Decode(&part)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to execute %s", op), zap.Error(err))
		return &model.Part{}, fmt.Errorf("%w: failed to execute %s: %s",
			serviceErrors.ErrInternalServer, op, err.Error())

	}

	return converter.ToModelPart(part), nil
}

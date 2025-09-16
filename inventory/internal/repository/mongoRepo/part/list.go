package part

import (
	"context"
	"fmt"
	"log"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/converter"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (r *repository) List(ctx context.Context) ([]*model.Part, error) {
	const op = "List"
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to execute %s", op), zap.Error(err))
		return []*model.Part{},
			fmt.Errorf("%w: failed to execute %s: %s", serviceErrors.ErrInternalServer, op, err.Error())
	}
	defer func() {
		cerr := cursor.Close(ctx)
		if cerr != nil {
			logger.Error(ctx, "failed to close cursor", zap.Error(cerr))
			log.Printf("failed to close cursor: %v\n", cerr.Error())
		}
	}()

	var parts []*repoModel.Part
	err = cursor.All(ctx, &parts)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to execute %s", op), zap.Error(err))
		return []*model.Part{},
			fmt.Errorf("%w: failed to execute %s: %s", serviceErrors.ErrInternalServer, op, err.Error())
	}

	return converter.ToModelParts(parts), nil
}

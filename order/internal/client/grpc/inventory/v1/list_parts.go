package v1

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/client/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"go.uber.org/zap"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	filterProto, err := converter.ToProtoPartsFilter(filter)
	if err != nil {
		logger.Error(ctx, "failed to get filter to get list of parts from inventory service", zap.Error(err))
		return []model.Part{}, err
	}

	req := &inventoryV1.ListPartsRequest{Filter: filterProto}
	resp, err := c.generatedClient.ListParts(ctx, req)
	if err != nil {
		logger.Error(ctx, "failed to get list of parts from inventory service", zap.Error(err))
		return []model.Part{}, err
	}

	if len(resp.Parts) == 0 {
		logger.Error(ctx, "empty list of parts from inventory service", zap.Error(err))
		return []model.Part{}, serviceErrors.ErrNotFoundFromRemoteInventory
	}

	return converter.ToModelParts(resp.Parts), nil
}

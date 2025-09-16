package v1

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/client/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	filterProto, err := converter.ToProtoPartsFilter(filter)
	if err != nil {
		return []model.Part{}, err
	}

	req := &inventoryV1.ListPartsRequest{Filter: filterProto}
	resp, err := c.generatedClient.ListParts(ctx, req)
	if err != nil {
		return []model.Part{}, err
	}

	if len(resp.Parts) == 0 {
		return []model.Part{}, serviceErrors.ErrNotFoundFromRemoteInventory
	}

	return converter.ToModelParts(resp.Parts), nil
}

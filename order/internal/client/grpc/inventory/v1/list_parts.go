package v1

import (
	"context"
	"github.com/kirillmc/starShipsCompany/order/internal/client/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func (c *client) ListParts(ctx context.Context, filter model.PartsFilter) ([]*model.Part, error) {
	filterProto := converter.PartsFilterToProto(filter)

	req := &inventoryV1.ListPartsRequest{Filter: &filterProto}
	resp, err := c.generatedClient.ListParts(ctx, req)
	if err != nil {
		return nil, err
	}

	parts := converter.PartsToModel(resp.Parts)

	return parts, nil
}

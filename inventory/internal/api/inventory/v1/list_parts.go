package v1

import (
	"context"

	"github.com/kirillmc/starShipsCompany/inventory/internal/converter"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	partsFilter := converter.ToModelPartsFilter(req.GetFilter())
	parts := a.inventoryService.List(ctx, partsFilter)

	return &inventoryV1.ListPartsResponse{
		Parts: converter.ToProtoParts(parts),
	}, nil
}

package v1

import (
	"context"
	"github.com/kirillmc/starShipsCompany/inventory/internal/converter"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func (s *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	partsFilter := converter.PartsFilterToModel(req.GetFilter())
	parts := s.inventoryService.List(ctx, partsFilter)

	partsProto := converter.PartsToProto(parts)

	return &inventoryV1.ListPartsResponse{
		Parts: partsProto,
	}, nil
}

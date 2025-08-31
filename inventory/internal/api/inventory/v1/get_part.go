package v1

import (
	"context"

	"github.com/kirillmc/starShipsCompany/inventory/internal/converter"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.Get(ctx, req.GetUuid())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to find part: %s", err)
	}

	partProto := converter.PartToProto(part)

	return &inventoryV1.GetPartResponse{
		Part: partProto,
	}, nil
}

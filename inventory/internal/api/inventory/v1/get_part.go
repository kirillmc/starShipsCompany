package v1

import (
	"context"
	"errors"

	"github.com/kirillmc/starShipsCompany/inventory/internal/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) GetPart(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error) {
	part, err := a.inventoryService.Get(ctx, req.GetUuid())
	if err != nil {
		if errors.Is(err, serviceErrors.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "failed to find part: %s", serviceErrors.ErrNotFound)
		}

		return nil, status.Errorf(codes.Internal, "failed to find part: %s", serviceErrors.ErrInternalServer)
	}

	return &inventoryV1.GetPartResponse{
		Part: converter.ToProtoPart(part),
	}, nil
}

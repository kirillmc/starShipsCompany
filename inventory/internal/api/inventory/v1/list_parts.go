package v1

import (
	"context"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"

	"github.com/kirillmc/starShipsCompany/inventory/internal/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error) {
	partsFilter := converter.ToModelPartsFilter(req.GetFilter())
	parts, err := a.inventoryService.List(ctx, partsFilter)
	if err != nil {
		logger.Error(ctx, "failed to find parts", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to find parts: %s", serviceErrors.ErrInternalServer)
	}

	return &inventoryV1.ListPartsResponse{
		Parts: converter.ToProtoParts(parts),
	}, nil
}

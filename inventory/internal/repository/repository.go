package repository

import (
	"context"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

type Repository interface {
	Get(ctx context.Context, req *inventoryV1.GetPartRequest) (*inventoryV1.GetPartResponse, error)
	List(ctx context.Context, req *inventoryV1.ListPartsRequest) (*inventoryV1.ListPartsResponse, error)
}

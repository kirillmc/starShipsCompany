package v1

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/service"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

type api struct {
	inventoryV1.UnimplementedInventoryServiceServer

	inventoryService service.Service
}

func NewAPI(inventoryService service.Service) *api {
	return &api{
		inventoryService: inventoryService,
	}
}

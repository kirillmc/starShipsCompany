package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToProtoManufacturer(manufacturer *model.Manufacturer) *inventoryV1.Manufacturer {
	manufacturerMapped := &inventoryV1.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}

	return manufacturerMapped
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToModelManufacturer(manufacturer *inventoryV1.Manufacturer) *model.Manufacturer {
	manufacturerMapped := &model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
	return manufacturerMapped
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ManufacturerToModel(manufacturer *inventoryV1.Manufacturer) *model.Manufacturer {
	manufacturerModel := &model.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}

	return manufacturerModel
}

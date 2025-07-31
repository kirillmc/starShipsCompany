package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func ManufacturerToService(manufacturer *model.Manufacturer) *serviceModel.Manufacturer {
	manufacturerService := &serviceModel.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}

	return manufacturerService
}

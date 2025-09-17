package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/model"
)

func ToModelManufacturer(manufacturer *model.Manufacturer) *serviceModel.Manufacturer {
	manufacturerMapped := &serviceModel.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}

	return manufacturerMapped
}

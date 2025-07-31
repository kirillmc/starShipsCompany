package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func CategoryToServiceModel(category model.Category) serviceModel.Category {
	switch category {
	case model.ENGINE:
		return serviceModel.ENGINE
	case model.FUEL:
		return serviceModel.FUEL
	case model.PORTHOLE:
		return serviceModel.PORTHOLE
	case model.WING:
		return serviceModel.WING
	default:
		return serviceModel.UNSPECIFIED
	}
}

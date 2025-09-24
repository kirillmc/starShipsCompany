package converter

import (
	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/model"
)

func ToModelCategory(category repoModel.Category) model.Category {
	switch category {
	case repoModel.Engine:
		return model.Engine
	case repoModel.Fuel:
		return model.Fuel
	case repoModel.Porthole:
		return model.Porthole
	case repoModel.Wing:
		return model.Wing
	default:
		return model.Unspecified
	}
}

package converter

import (
	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/model"
)

func ToModelDimensions(dimensions *repoModel.Dimensions) *model.Dimensions {
	dimensionsMapped := &model.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}

	return dimensionsMapped
}

package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func DimensionsToService(dimensions *model.Dimensions) *serviceModel.Dimensions {
	dimensionsService := &serviceModel.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}

	return dimensionsService
}

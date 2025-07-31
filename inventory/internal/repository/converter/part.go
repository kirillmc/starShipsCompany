package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"

	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func PartToServiceModel(part *model.Part) *serviceModel.Part {
	partService := &serviceModel.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      CategoryToServiceModel(part.Category),
		Dimensions:    DimensionsToServiceModel(part.Dimensions),
		Manufacturer:  part.Manufacturer,
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}

	return &serviceModel.Part{}
}

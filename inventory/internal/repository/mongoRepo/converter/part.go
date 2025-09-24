package converter

import (
	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo/model"
)

func ToModelPart(part *repoModel.Part) *model.Part {
	partMapped := &model.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      ToModelCategory(part.Category),
		Dimensions:    ToModelDimensions(part.Dimensions),
		Manufacturer:  ToModelManufacturer(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      part.Metadata,
		CreatedAt:     part.CreatedAt,
		UpdatedAt:     part.UpdatedAt,
	}

	return partMapped
}

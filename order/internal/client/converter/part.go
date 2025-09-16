package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
)

func ToModelPart(part *inventoryV1.Part) model.Part {
	if part == nil {
		return model.Part{}
	}

	partMapped := model.Part{
		UUID:          part.Uuid,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      ToModelCategory(part.Category),
		Dimensions:    ToModelDimensions(part.Dimensions),
		Manufacturer:  ToModelManufacturer(part.Manufacturer),
		Tags:          part.Tags,
		Metadata:      ToModelMetadata(part.Metadata),
		CreatedAt:     lo.ToPtr(part.CreatedAt.AsTime()),
		UpdatedAt:     lo.ToPtr(part.UpdatedAt.AsTime()),
	}

	return partMapped
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToModelCategories(categories []inventoryV1.CATEGORY) []model.Category {
	categoriesMapped := make([]model.Category, 0, len(categories))
	for _, category := range categories {
		categoriesMapped = append(categoriesMapped, ToModelCategory(category))
	}
	return categoriesMapped
}

func ToModelCategory(category inventoryV1.CATEGORY) model.Category {
	switch category {
	case inventoryV1.CATEGORY_ENGINE:
		return model.Engine
	case inventoryV1.CATEGORY_FUEL:
		return model.Fuel
	case inventoryV1.CATEGORY_PORTHOLE:
		return model.Porthole
	case inventoryV1.CATEGORY_WING:
		return model.Wing
	default:
		return model.Unspecified
	}
}

func ToProtoCategory(category model.Category) inventoryV1.CATEGORY {
	switch category {
	case model.Engine:
		return inventoryV1.CATEGORY_ENGINE
	case model.Fuel:
		return inventoryV1.CATEGORY_FUEL
	case model.Porthole:
		return inventoryV1.CATEGORY_PORTHOLE
	case model.Wing:
		return inventoryV1.CATEGORY_WING
	default:
		return inventoryV1.CATEGORY_UNSPECIFIED
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToProtoCategories(categories []model.Category) []inventoryV1.CATEGORY {
	categoriesMapped := make([]inventoryV1.CATEGORY, 0, len(categories))
	for _, category := range categories {
		categoriesMapped = append(categoriesMapped, ToProtoCategory(category))
	}
	return categoriesMapped
}

func ToProtoCategory(category model.Category) inventoryV1.CATEGORY {
	switch category {
	case model.CategoryEngine:
		return inventoryV1.CATEGORY_ENGINE
	case model.CategoryFuel:
		return inventoryV1.CATEGORY_FUEL
	case model.CategoryPorthole:
		return inventoryV1.CATEGORY_PORTHOLE
	case model.CategoryWing:
		return inventoryV1.CATEGORY_WING
	default:
		return inventoryV1.CATEGORY_UNSPECIFIED
	}
}

func ToModelCategory(category inventoryV1.CATEGORY) model.Category {
	switch category {
	case inventoryV1.CATEGORY_ENGINE:
		return model.CategoryEngine
	case inventoryV1.CATEGORY_FUEL:
		return model.CategoryFuel
	case inventoryV1.CATEGORY_PORTHOLE:
		return model.CategoryPorthole
	case inventoryV1.CATEGORY_WING:
		return model.CategoryWing
	default:
		return model.CategoryUnspecified
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func CategoriesToProto(categories []model.Category) []inventoryV1.CATEGORY {
	categoriesProto := make([]inventoryV1.CATEGORY, 0, len(categories))

	for _, category := range categories {
		categoriesProto = append(categoriesProto, CategoryToProto(category))
	}

	return categoriesProto
}

func CategoryToProto(category model.Category) inventoryV1.CATEGORY {
	switch category {
	case model.ENGINE:
		return inventoryV1.CATEGORY_ENGINE
	case model.FUEL:
		return inventoryV1.CATEGORY_FUEL
	case model.PORTHOLE:
		return inventoryV1.CATEGORY_PORTHOLE
	case model.WING:
		return inventoryV1.CATEGORY_WING
	default:
		return inventoryV1.CATEGORY_UNSPECIFIED
	}
}

func CategoryToModel(category inventoryV1.CATEGORY) model.Category {
	switch category {
	case inventoryV1.CATEGORY_ENGINE:
		return model.ENGINE
	case inventoryV1.CATEGORY_FUEL:
		return model.FUEL
	case inventoryV1.CATEGORY_PORTHOLE:
		return model.PORTHOLE
	case inventoryV1.CATEGORY_WING:
		return model.WING
	default:
		return model.UNSPECIFIED_CATEGORY
	}
}

package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func PartsFilterToProto(partsFiler *model.PartsFilter) *inventoryV1.PartsFilter {
	partsFilter := &inventoryV1.PartsFilter{
		Uuids:                 partsFiler.UUIDs,
		Names:                 partsFiler.Names,
		Categories:            CategoriesToProto(partsFiler.Categories),
		ManufacturerCountries: partsFiler.ManufacturerCountries,
		Tags:                  partsFiler.Tags,
	}

	return partsFilter
}

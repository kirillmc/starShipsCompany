package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToModelPartsFilter(partsFiler *inventoryV1.PartsFilter) *model.PartsFilter {
	partsFilterMapped := &model.PartsFilter{
		UUIDs:                 partsFiler.Uuids,
		Names:                 partsFiler.Names,
		Categories:            ToModelCategories(partsFiler.Categories),
		ManufacturerCountries: partsFiler.ManufacturerCountries,
		Tags:                  partsFiler.Tags,
	}

	return partsFilterMapped
}

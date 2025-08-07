package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func PartsFilterToModel(partsFiler *inventoryV1.PartsFilter) *model.PartFilter {
	partsFiletModel := &model.PartFilter{
		UUIDs:                 partsFiler.Uuids,
		Names:                 partsFiler.Names,
		Categories:            CATEGORIESToModel(partsFiler.Categories),
		ManufacturerCountries: partsFiler.ManufacturerCountries,
		Tags:                  partsFiler.Tags,
	}

	return partsFiletModel
}

package converter

import (
	"fmt"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func PartsFilterToProto(partsFiler model.PartsFilter) (*inventoryV1.PartsFilter, error) {
	err := checkFilterParams(partsFiler)
	if err != nil {
		return &inventoryV1.PartsFilter{}, err
	}

	partsFilter := &inventoryV1.PartsFilter{
		Uuids:                 partsFiler.UUIDs,
		Names:                 partsFiler.Names,
		Categories:            CategoriesToProto(partsFiler.Categories),
		ManufacturerCountries: partsFiler.ManufacturerCountries,
		Tags:                  partsFiler.Tags,
	}

	return partsFilter, nil
}

func checkFilterParams(partsFiler model.PartsFilter) error {
	if len(partsFiler.UUIDs) == 0 && len(partsFiler.Names) == 0 && len(partsFiler.Categories) == 0 &&
		len(partsFiler.ManufacturerCountries) == 0 && len(partsFiler.Tags) == 0 {
		return fmt.Errorf("filter params is empty")
	}

	return nil
}

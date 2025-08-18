package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func PartsToModel(parts []*inventoryV1.Part) []*model.Part {
	partsModel := make([]*model.Part, 0, len(parts))

	for _, part := range parts {
		partsModel = append(partsModel, PartToModel(part))
	}

	return partsModel
}

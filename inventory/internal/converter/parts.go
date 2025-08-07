package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func PartsToProto(parts []*model.Part) []*inventoryV1.Part {
	partsProto := make([]*inventoryV1.Part, 0, len(parts))

	for _, part := range parts {
		partsProto = append(partsProto, PartToProto(part))
	}

	return partsProto
}

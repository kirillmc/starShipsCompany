package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToModelParts(parts []*inventoryV1.Part) []model.Part {
	partsMapped := make([]model.Part, 0, len(parts))
	for _, part := range parts {
		partsMapped = append(partsMapped, ToModelPart(part))
	}
	return partsMapped
}

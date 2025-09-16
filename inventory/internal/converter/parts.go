package converter

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
)

func ToProtoParts(parts []*model.Part) []*inventoryV1.Part {
	partsMapped := make([]*inventoryV1.Part, 0, len(parts))
	for _, part := range parts {
		partsMapped = append(partsMapped, ToProtoPart(part))
	}

	return partsMapped
}

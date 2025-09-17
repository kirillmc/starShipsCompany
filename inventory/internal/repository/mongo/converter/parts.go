package converter

import (
	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/model"
)

func ToModelParts(parts []*repoModel.Part) []*model.Part {
	partsMapped := make([]*model.Part, 0, len(parts))
	for _, part := range parts {
		partsMapped = append(partsMapped, ToModelPart(part))
	}

	return partsMapped
}

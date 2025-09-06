package converter

import (
	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func ToModelParts(parts map[string]*repoModel.Part) map[string]*model.Part {
	partsMapped := make(map[string]*model.Part, len(parts))
	for uuid, part := range parts {
		partsMapped[uuid] = ToModelPart(part)
	}

	return partsMapped
}

package converter

import (
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func PartsToService(parts map[string]*model.Part) map[string]*serviceModel.Part {
	partsService := make(map[string]*serviceModel.Part, len(parts))
	for uuid, part := range parts {
		partsService[uuid] = PartToService(part)
	}

	return partsService
}

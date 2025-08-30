package part

import (
	"context"

	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func (r *repository) List(_ context.Context) map[model.PartUUID]*serviceModel.Part {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return converter.PartsToService(r.parts)
}

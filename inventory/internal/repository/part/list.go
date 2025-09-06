package part

import (
	"context"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/converter"
)

func (r *repository) List(_ context.Context) map[model.PartUUID]*model.Part {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return converter.ToModelParts(r.parts)
}

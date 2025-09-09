package part

import (
	"context"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/converter"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func (r *repository) List(_ context.Context) map[model.PartUUID]*model.Part {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return converter.ToModelParts(r.parts)
}

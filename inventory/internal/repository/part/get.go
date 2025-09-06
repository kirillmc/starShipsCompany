package part

import (
	"context"
	"fmt"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/serviceErrors"
)

func (r *repository) Get(_ context.Context, partUUID model.PartUUID) (*model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	const op = "Get"

	part, ok := r.parts[partUUID]
	if !ok {
		return &model.Part{}, fmt.Errorf("failed to execute %s method: failed to get part: %w",
			op, serviceErrors.ErrNotFound)
	}

	return converter.ToModelPart(part), nil
}

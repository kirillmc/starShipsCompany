package part

import (
	"context"
	"fmt"
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func (r *repository) Get(_ context.Context, uuid model.UUID) (*serviceModel.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	const op = "Get"

	part, ok := r.parts[uuid]
	if !ok {
		return nil, fmt.Errorf("failed to execute %s method: part with UUID %s not found", op, uuid)
	}

	return converter.PartToService(part), nil
}

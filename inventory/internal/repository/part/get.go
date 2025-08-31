package part

import (
	"context"
	"fmt"

	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/converter"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

func (r *repository) Get(_ context.Context, partUUID model.PartUUID) (*serviceModel.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	const op = "Get"

	part, ok := r.parts[partUUID]
	if !ok {
		return &serviceModel.Part{}, fmt.Errorf("failed to execute %s method: part with PartUUID %s not found",
			op, partUUID)
	}

	return converter.PartToService(part), nil
}

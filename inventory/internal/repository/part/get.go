package part

import (
	"context"
	"fmt"
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/converter"
)

func (r *repository) GetPart(_ context.Context, uuid string) (*serviceModel.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[uuid]
	if !ok {
		return nil, fmt.Errorf("part with UUID %r not found", uuid)
	}

	return converter.PartToServiceModel(part), nil
}

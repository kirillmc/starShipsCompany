package part

import (
	"context"

	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func (s *service) Get(ctx context.Context, partUUID model.PartUUID) (*model.Part, error) {
	part, err := s.repo.Get(ctx, partUUID)
	if err != nil {
		return &model.Part{}, err
	}

	return part, nil
}

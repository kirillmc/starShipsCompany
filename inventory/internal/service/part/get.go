package part

import (
	"context"
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

func (s *service) Get(ctx context.Context, uuid model.UUID) (*model.Part, error) {
	part, err := s.repo.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return part, nil
}

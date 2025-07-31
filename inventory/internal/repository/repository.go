package repository

import (
	"context"
	serviceModel "github.com/kirillmc/starShipsCompany/inventory/internal/model"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

type Repository interface {
	Get(ctx context.Context, uuid model.UUID) (*serviceModel.Part, error)
	List(ctx context.Context) map[model.UUID]*serviceModel.Part
}

package mongo

import (
	"context"

	model "github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

type Repository interface {
	Get(ctx context.Context, uuid model.PartUUID) (*model.Part, error)
	List(ctx context.Context) ([]*model.Part, error)
}

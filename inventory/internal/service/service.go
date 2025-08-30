package service

import (
	"context"
	"github.com/kirillmc/starShipsCompany/inventory/internal/model"
)

type Service interface {
	Get(ctx context.Context, uuid model.PartUUID) (*model.Part, error)
	List(ctx context.Context, filter *model.PartsFilter) []*model.Part
}

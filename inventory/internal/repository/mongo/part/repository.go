package part

import (
	def "github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/model"
	"sync"
)

var _ def.Repository = (*repository)(nil)

type repository struct {
	mu    sync.RWMutex
	parts map[string]*model.Part
}

func NewRepository() *repository {
	return &repository{parts: setDefaultPartsMap()}
}

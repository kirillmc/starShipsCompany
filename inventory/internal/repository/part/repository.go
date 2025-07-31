package part

import (
	def "github.com/kirillmc/starShipsCompany/inventory/internal/repository"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
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

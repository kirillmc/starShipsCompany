package part

import (
	"sync"

	def "github.com/kirillmc/starShipsCompany/inventory/internal/repository"
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/model"
)

var _ def.Repository = (*repository)(nil)

type repository struct {
	mu    sync.RWMutex
	parts map[string]*model.Part
}

func NewRepository() *repository {
	return &repository{parts: setDefaultPartsMap()}
}

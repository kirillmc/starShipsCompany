package order

import (
	"sync"

	def "github.com/kirillmc/starShipsCompany/order/internal/repository"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

var _ def.Repository = (*repository)(nil)

type repository struct {
	mu     sync.RWMutex
	orders map[string]*model.Order
}

func NewRepository() *repository {
	return &repository{
		orders: make(map[string]*model.Order),
	}
}

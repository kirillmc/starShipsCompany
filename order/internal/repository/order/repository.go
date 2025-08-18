package order

import (
	def "github.com/kirillmc/starShipsCompany/order/internal/repository"
	"sync"
)

var _ def.Repository = (*repository)(nil)

type repository struct {
	mu     sync.RWMutex
	orders map[string]*Order
}

func NewOrderStorage() *repository {
	return &repository{
		orders: make(map[string]*Order),
	}
}

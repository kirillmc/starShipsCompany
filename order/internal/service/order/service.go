package order

import (
	"github.com/kirillmc/starShipsCompany/order/internal/repository"
	def "github.com/kirillmc/starShipsCompany/order/internal/service"
)

var _ def.Service = (*service)(nil)

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

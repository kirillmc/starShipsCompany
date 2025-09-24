package part

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongoRepo"
	def "github.com/kirillmc/starShipsCompany/inventory/internal/service"
)

var _ def.Service = (*service)(nil)

type service struct {
	repo mongoRepo.Repository
}

func NewService(repo mongoRepo.Repository) *service {
	return &service{
		repo: repo,
	}
}

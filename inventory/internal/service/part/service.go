package part

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo"
	def "github.com/kirillmc/starShipsCompany/inventory/internal/service"
)

var _ def.Service = (*service)(nil)

type service struct {
	repo mongo.Repository
}

func NewService(repo mongo.Repository) *service {
	return &service{
		repo: repo,
	}
}

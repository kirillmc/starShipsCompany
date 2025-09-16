package payment

import def "github.com/kirillmc/starShipsCompany/payment/internal/service"

var _ def.Service = (*service)(nil)

type service struct{}

func NewService() *service {
	return &service{}
}

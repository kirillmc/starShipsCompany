package v1

import (
	"context"
	"testing"

	"github.com/kirillmc/starShipsCompany/inventory/internal/service/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	service *mocks.Service

	api *api
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.service = mocks.NewService(s.T())

	s.api = NewAPI(
		s.service,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

package v1

import (
	"testing"

	"github.com/kirillmc/starShipsCompany/inventory/internal/service/mocks"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	service *mocks.Service

	api *api
}

func (s *ServiceSuite) SetupTest() {
	s.service = mocks.NewService(s.T())

	err := logger.Init("", true)
	s.Require().NoError(err)

	s.api = NewAPI(
		s.service,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

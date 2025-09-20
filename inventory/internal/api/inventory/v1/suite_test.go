package v1

import (
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"testing"

	"github.com/kirillmc/starShipsCompany/inventory/internal/service/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	service *mocks.Service

	api *api
}

func (s *ServiceSuite) SetupTest() {
	s.service = mocks.NewService(s.T())

	logger.Init("", true)

	s.api = NewAPI(
		s.service,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

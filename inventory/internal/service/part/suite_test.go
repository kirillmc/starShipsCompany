package part

import (
	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mongo/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	repository *mocks.Repository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.repository = mocks.NewRepository(s.T())

	s.service = NewService(
		s.repository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

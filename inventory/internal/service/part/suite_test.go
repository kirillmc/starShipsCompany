package part

import (
	"context"
	"testing"

	"github.com/kirillmc/starShipsCompany/inventory/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repository *mocks.Repository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

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

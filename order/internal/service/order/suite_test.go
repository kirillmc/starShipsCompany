package order

import (
	"context"
	"testing"

	clienMocks "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/mocks"
	repoMock "github.com/kirillmc/starShipsCompany/order/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	repository      *repoMock.Repository
	inventoryClient *clienMocks.InventoryClient
	paymentClient   *clienMocks.PaymentClient

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.inventoryClient = clienMocks.NewInventoryClient(s.T())
	s.paymentClient = clienMocks.NewPaymentClient(s.T())
	s.repository = repoMock.NewRepository(s.T())

	s.service = NewService(
		s.inventoryClient,
		s.paymentClient,
		s.repository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

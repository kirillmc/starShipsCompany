package order

import (
	repoMock "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/mocks"
	"testing"

	clienMocks "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	repository      *repoMock.Repository
	inventoryClient *clienMocks.InventoryClient
	paymentClient   *clienMocks.PaymentClient

	service *service
}

func (s *ServiceSuite) SetupTest() {
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

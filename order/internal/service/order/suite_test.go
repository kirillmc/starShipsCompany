package order

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	clienMocks "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/mocks"
	repoMock "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	orderRepository     *repoMock.OrderRepository
	orderPartRepository *repoMock.OrderPartRepository
	inventoryClient     *clienMocks.InventoryClient
	paymentClient       *clienMocks.PaymentClient

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.inventoryClient = clienMocks.NewInventoryClient(s.T())
	s.paymentClient = clienMocks.NewPaymentClient(s.T())
	s.orderRepository = repoMock.NewOrderRepository(s.T())
	s.orderPartRepository = repoMock.NewOrderPartRepository(s.T())

	pool := &pgxpool.Pool{}
	s.service = NewService(
		pool,
		s.inventoryClient,
		s.paymentClient,
		s.orderRepository,
		s.orderPartRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

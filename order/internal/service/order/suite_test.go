package order

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	clienMocks "github.com/kirillmc/starShipsCompany/order/internal/client/grpc/mocks"
	repoMocks "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/mocks"
	serviceMocks "github.com/kirillmc/starShipsCompany/order/internal/service/mocks"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	orderRepository     *repoMocks.OrderRepository
	orderPartRepository *repoMocks.OrderPartRepository
	inventoryClient     *clienMocks.InventoryClient
	paymentClient       *clienMocks.PaymentClient
	orderPaidProducer   *serviceMocks.OrderProducerService

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.inventoryClient = clienMocks.NewInventoryClient(s.T())
	s.paymentClient = clienMocks.NewPaymentClient(s.T())
	s.orderPaidProducer = serviceMocks.NewOrderProducerService(s.T())
	s.orderRepository = repoMocks.NewOrderRepository(s.T())
	s.orderPartRepository = repoMocks.NewOrderPartRepository(s.T())

	err := logger.Init("", true)
	s.Require().NoError(err)

	pool := &pgxpool.Pool{}
	s.service = NewService(
		pool,
		s.inventoryClient,
		s.paymentClient,
		s.orderPaidProducer,
		s.orderRepository,
	)
}

func (s *ServiceSuite) TearDownTest() {
}

func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

package orderConsumer

import (
	"context"
	kafkaConverter "github.com/kirillmc/starShipsCompany/assembly/internal/converter/kafka"
	def "github.com/kirillmc/starShipsCompany/assembly/internal/service"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.ConsumerService = (*service)(nil)

type service struct {
	orderPaidConsumer kafka.Consumer
	orderPaidDecoder  kafkaConverter.OrderPaidDecoder

	orderAssembledProducer def.OrderProducerService
}

func NewService(orderPaidConsumer kafka.Consumer,
	orderPaidDecoder kafkaConverter.OrderPaidDecoder,
	orderAssembledProducer def.OrderProducerService) *service {
	return &service{
		orderPaidConsumer:      orderPaidConsumer,
		orderPaidDecoder:       orderPaidDecoder,
		orderAssembledProducer: orderAssembledProducer,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order paid consumer service")

	err := s.orderPaidConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.paid topic error", zap.Error(err))
		return err
	}

	return nil
}

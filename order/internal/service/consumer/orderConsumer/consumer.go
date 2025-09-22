package orderConsumer

import (
	"context"
	kafkaConverter "github.com/kirillmc/starShipsCompany/order/internal/converter/kafka"
	def "github.com/kirillmc/starShipsCompany/order/internal/service"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.ConsumerService = (*service)(nil)

type service struct {
	orderAssembledConsumer kafka.Consumer
	orderAssembledDecoder  kafkaConverter.Decoder

	orderService def.Service
}

func NewService(orderAssembledConsumer kafka.Consumer,
	orderAssembledDecoder kafkaConverter.Decoder,
	orderService def.Service) *service {
	return &service{
		orderAssembledConsumer: orderAssembledConsumer,
		orderAssembledDecoder:  orderAssembledDecoder,
		orderService:           orderService,
	}
}

func (s *service) RunConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order paid consumer service")

	err := s.orderAssembledConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.paid topic error", zap.Error(err))
		return err
	}

	return nil
}

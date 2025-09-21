package orderAssembledConsumer

import (
	"context"
	kafkaConverter "github.com/kirillmc/starShipsCompany/notification/internal/converter/kafka"
	def "github.com/kirillmc/starShipsCompany/notification/internal/service"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.ConsumerService = (*service)(nil)

type service struct {
	orderAssembledConsumer kafka.Consumer
	orderAssembledDecoder  kafkaConverter.OrderAssembledDecoder
}

func NewService(orderAssembledConsumer kafka.Consumer,
	orderAssembledDecoder kafkaConverter.OrderAssembledDecoder) *service {
	return &service{
		orderAssembledConsumer: orderAssembledConsumer,
		orderAssembledDecoder:  orderAssembledDecoder,
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

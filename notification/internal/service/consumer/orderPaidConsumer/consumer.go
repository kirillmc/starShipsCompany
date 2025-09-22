package orderPaidConsumer

import (
	"context"
	kafkaConverter "github.com/kirillmc/starShipsCompany/notification/internal/converter/kafka"
	def "github.com/kirillmc/starShipsCompany/notification/internal/service"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.OrderPaidConsumerService = (*service)(nil)

type service struct {
	orderPaidConsumer kafka.Consumer
	orderPaidDecoder  kafkaConverter.OrderPaidDecoder

	telegramService def.TelegramService
}

func NewService(orderPaidConsumer kafka.Consumer,
	orderPaidDecoder kafkaConverter.OrderPaidDecoder,
	telegramService def.TelegramService) *service {
	return &service{
		orderPaidConsumer: orderPaidConsumer,
		orderPaidDecoder:  orderPaidDecoder,

		telegramService: telegramService,
	}
}

func (s *service) RunOrderPaidConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order paid consumer service")

	err := s.orderPaidConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.paid topic error", zap.Error(err))
		return err
	}

	return nil
}

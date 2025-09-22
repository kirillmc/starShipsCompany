package orderAssembledConsumer

import (
	"context"

	kafkaConverter "github.com/kirillmc/starShipsCompany/notification/internal/converter/kafka"
	def "github.com/kirillmc/starShipsCompany/notification/internal/service"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

var _ def.OrderAssembledConsumerService = (*service)(nil)

type service struct {
	orderAssembledConsumer kafka.Consumer
	orderAssembledDecoder  kafkaConverter.OrderAssembledDecoder

	telegramService def.TelegramService
}

func NewService(orderAssembledConsumer kafka.Consumer,
	orderAssembledDecoder kafkaConverter.OrderAssembledDecoder,
	telegramService def.TelegramService,
) *service {
	return &service{
		orderAssembledConsumer: orderAssembledConsumer,
		orderAssembledDecoder:  orderAssembledDecoder,
		telegramService:        telegramService,
	}
}

func (s *service) RunOrderAssembledConsumer(ctx context.Context) error {
	logger.Info(ctx, "Starting order assembler consumer service")

	err := s.orderAssembledConsumer.Consume(ctx, s.OrderHandler)
	if err != nil {
		logger.Error(ctx, "Consume from order.assembler topic error", zap.Error(err))
		return err
	}

	return nil
}

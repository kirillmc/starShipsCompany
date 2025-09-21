package orderAssembledConsumer

import (
	"context"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka/consumer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

func (s *service) OrderHandler(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderAssembledDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderAssembled", zap.Error(err))
		return err
	}

	_ = event

	return nil
}

package orderConsumer

import (
	"context"
	"time"

	"github.com/kirillmc/starShipsCompany/assembly/internal/converter"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka/consumer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

const secondsToBuildShip = 10

func (s *service) OrderHandler(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderPaidDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderPaid", zap.Error(err))
		return err
	}

	time.Sleep(secondsToBuildShip * time.Second) //nolint:forbidigo // здесь нужен sleep для симуляции

	err = s.orderAssembledProducer.ProduceOrderAssembled(ctx, converter.ToEventOrderAssembled(event, secondsToBuildShip))
	if err != nil {
		logger.Error(ctx, "Failed to produce OrderAssembled", zap.Error(err))
		return err
	}

	return nil
}

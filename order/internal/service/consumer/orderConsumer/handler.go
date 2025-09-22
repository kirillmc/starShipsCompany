package orderConsumer

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka/consumer"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (s *service) OrderHandler(ctx context.Context, msg consumer.Message) error {
	event, err := s.orderAssembledDecoder.Decode(msg.Value)
	if err != nil {
		logger.Error(ctx, "Failed to decode OrderAssembled", zap.Error(err))
		return err
	}

	err = s.orderService.Update(ctx,
		model.UpdateOrderParams{
			OrderUUID: event.OrderUUID,
			Status:    lo.ToPtr(model.OrderStatusAssembled),
		})
	if err != nil {
		logger.Error(ctx, "Failed to set order status ASSEMBLED", zap.Error(err))
		return err
	}

	return nil
}

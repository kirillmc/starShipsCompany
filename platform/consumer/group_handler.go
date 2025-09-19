package consumer

import (
	"context"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// MessageHandler — обработчик сообщений.
type MessageHandler func(ctx context.Context, msg Message) error

// Middleware — функция middleware для дополнительной обработки.
type Middleware func(next MessageHandler) MessageHandler

// groupHandler — обёртка для sarama.ConsumerGroupHandler
type groupHandler struct {
	handler MessageHandler
	logger  Logger
}

// NewGroupHandler создаёт новый groupHandler с middleware цепочкой.
func NewGroupHandler(handler MessageHandler, logger Logger, middlewares ...Middleware) *groupHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		// Применяем middleware цепочку
		handler = middlewares[i](handler)
	}

	return &groupHandler{
		handler: handler,
		logger:  logger,
	}
}

func (g *groupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (g *groupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (g *groupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				g.logger.Info(session.Context(), "Kafka message channel closed")
				return nil
			}

			msg := Message{
				Headers:        extractHeaders(message.Headers),
				Timestamp:      message.Timestamp,
				BlockTimestamp: message.BlockTimestamp,
				Key:            message.Key,
				Value:          message.Value,
				Topic:          message.Topic,
				Partition:      message.Partition,
				Offset:         message.Offset,
			}

			// TODO: nонять, что есть handler
			if err := g.handler(session.Context(), msg); err != nil {
				g.logger.Error(session.Context(), "Kafka handler error", zap.Error(err))
				continue
			}

			session.MarkMessage(message, "")

		case <-session.Context().Done():
			g.logger.Info(session.Context(), "Kafka session context done")
			return nil
		}
	}
}

func extractHeaders(headers []*sarama.RecordHeader) map[string][]byte {
	result := make(map[string][]byte)
	for _, header := range headers {
		if header != nil && len(header.Key) > 0 {
			result[string(header.Key)] = header.Value
		}
	}

	return result
}

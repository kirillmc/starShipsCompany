package orderProducer

import (
	"context"
	"github.com/kirillmc/starShipsCompany/assembly/internal/model"
	"github.com/kirillmc/starShipsCompany/platform/pkg/kafka"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	eventsV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/events/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type service struct {
	orderAssembledProducer kafka.Producer
}

func NewService(orderAssembledProducer kafka.Producer) *service {
	return &service{
		orderAssembledProducer: orderAssembledProducer,
	}
}

func (p *service) ProduceOrderAssembled(ctx context.Context, event model.OrderAssembledEvent) error {
	msg := &eventsV1.ShipAssembled{
		EventUuid:    event.EventUUID,
		OrderUuid:    event.OrderUUID,
		UserUuid:     event.UserUUID,
		BuildTimeSec: event.BuildTimeSec,
	}

	payload, err := proto.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "failed to marshal OrderAssembled", zap.Error(err))
		return err
	}

	err = p.orderAssembledProducer.Send(ctx, []byte(event.EventUUID), payload)
	if err != nil {
		logger.Error(ctx, "failed to publish OrderAssembled", zap.Error(err))
		return err
	}

	return nil
}

package service

import (
	"context"
	"github.com/kirillmc/starShipsCompany/assembly/internal/model"
)

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}

type OrderProducerService interface {
	ProduceOrderAssembled(ctx context.Context, event model.OrderAssembledEvent) error
}

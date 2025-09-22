package service

import (
	"context"

	"github.com/kirillmc/starShipsCompany/notification/internal/model"
)

type OrderPaidConsumerService interface {
	RunOrderPaidConsumer(ctx context.Context) error
}

type OrderAssembledConsumerService interface {
	RunOrderAssembledConsumer(ctx context.Context) error
}

type TelegramService interface {
	SendOrderAssembledNotification(ctx context.Context, orderAssembledInfo model.OrderAssembledEvent) error
	SendOrderPaidNotification(ctx context.Context, orderPaidInfo model.OrderPaidEvent) error
}

package service

import (
	"context"
	"github.com/kirillmc/starShipsCompany/payment/internal/model"
)

type Service interface {
	Pay(context.Context, *model.PayOrderInfo) model.UUID
}

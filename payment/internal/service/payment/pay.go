package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/payment/internal/model"
)

func (s *service) Pay(context.Context, *model.PayOrderInfo) model.TransactionUUID {
	transactionUUID := uuid.NewString()

	return transactionUUID
}

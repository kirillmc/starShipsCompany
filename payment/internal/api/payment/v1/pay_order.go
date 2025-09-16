package v1

import (
	"context"
	"log"

	"github.com/kirillmc/starShipsCompany/payment/internal/converter"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func (a *api) PayOrder(ctx context.Context, req *paymentV1.PayOrderRequest) (*paymentV1.PayOrderResponse, error) {
	transactionUUID := a.paymentService.Pay(ctx, converter.ToModelPayOrderInfo(req))
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", transactionUUID)

	return &paymentV1.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}

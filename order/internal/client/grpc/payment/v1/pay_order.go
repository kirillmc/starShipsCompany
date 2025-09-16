package v1

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/client/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/platform/pkg/logger"
	"go.uber.org/zap"
)

func (c *client) PayOrder(ctx context.Context, payOrderInfo model.PayOrderParams) (model.TransactionUUID, error) {
	resp, err := c.generatedClient.PayOrder(ctx, converter.ToProtoPayOrderRequest(payOrderInfo))
	if err != nil {
		logger.Error(ctx, "failed to pay for order to payment service", zap.Error(err))
		return "", err
	}

	return resp.TransactionUuid, nil
}

package v1

import (
	"context"

	"github.com/kirillmc/starShipsCompany/order/internal/client/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (c *client) PayOrder(ctx context.Context, payOrderInfo model.PayOrderParams) (model.TransactionUUID, error) {
	resp, err := c.generatedClient.PayOrder(ctx, converter.ToProtoPayOrderRequest(payOrderInfo))
	if err != nil {
		return "", err
	}

	return resp.TransactionUuid, nil
}

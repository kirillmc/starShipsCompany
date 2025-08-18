package v1

import (
	"context"
	"github.com/kirillmc/starShipsCompany/order/internal/client/converter"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
)

func (c *client) PayOrder(ctx context.Context, payOrderInfo model.PayOrderInfo) (model.UUID, error) {
	req := converter.ModelToPayOrderRequest(payOrderInfo)
	resp, err := c.generatedClient.PayOrder(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.TransactionUuid, nil
}

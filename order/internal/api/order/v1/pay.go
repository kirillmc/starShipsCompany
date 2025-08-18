package v1

import (
	"context"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
	"net/http"
)

func (a *api) PayOrder(ctx context.Context, req *orderV1.PayOrderRequest, params orderV1.PayOrderParams) (orderV1.PayOrderRes, error) {
	if req == nil {
		return &orderV1.PayOrderResponse{}, nil
	}

	order, err := a.storage.getOrder(params.OrderUUID.String())
	if err != nil {
		return &orderV1.NotFoundError{
			Code:    404,
			Message: err.Error(),
		}, nil
	}

	paymentReq := paymentV1.PayOrderRequest{
		OrderUuid:     params.OrderUUID.String(),
		UserUuid:      order.UserUUID,
		PaymentMethod: paymentMethodToPaymentV1(req.PaymentMethod),
	}

	resp, err := a.paymentService.PayOrder(ctx, &paymentReq)
	if err != nil || resp == nil {
		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}, nil
	}

	a.storage.setOrderStatus(order.OrderUUID, orderV1.OrderStatusPAID)

	return &orderV1.PayOrderResponse{TransactionUUID: resp.TransactionUuid}, nil
}

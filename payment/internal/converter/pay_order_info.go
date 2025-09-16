package converter

import (
	"github.com/kirillmc/starShipsCompany/payment/internal/model"
	paymentV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/payment/v1"
)

func ToModelPayOrderInfo(req *paymentV1.PayOrderRequest) *model.PayOrderInfo {
	if req == nil {
		return nil
	}

	payOrderInfo := model.PayOrderInfo{
		OrderUUID:     req.OrderUuid,
		UserUUID:      req.UserUuid,
		PaymentMethod: ToModelPaymentMethod(req.PaymentMethod),
	}

	return &payOrderInfo
}

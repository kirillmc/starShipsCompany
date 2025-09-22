package converter

import (
	uuid "github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/samber/lo"
)

func ToEventOrderPaid(orderParams model.UpdateOrderParams) model.OrderPaidEvent {
	eventMapped := model.OrderPaidEvent{
		EventUUID:       uuid.NewString(),
		OrderUUID:       orderParams.OrderUUID,
		UserUUID:        "TEST",
		PaymentMethod:   string(lo.FromPtrOr(orderParams.PaymentMethod, model.PaymentMethodUnspecified)),
		TransactionUUID: lo.FromPtr(orderParams.TransactionUUID),
	}

	return eventMapped
}

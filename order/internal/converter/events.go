package converter

import (
	uuid "github.com/google/uuid"
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/samber/lo"
)

const testUUID = "123e4567-e89b-12d3-a456-426614174000"

func ToEventOrderPaid(orderParams model.UpdateOrderParams) model.OrderPaidEvent {
	eventMapped := model.OrderPaidEvent{
		EventUUID:       uuid.NewString(),
		OrderUUID:       orderParams.OrderUUID,
		UserUUID:        testUUID,
		PaymentMethod:   string(lo.FromPtrOr(orderParams.PaymentMethod, model.PaymentMethodUnspecified)),
		TransactionUUID: lo.FromPtr(orderParams.TransactionUUID),
	}

	return eventMapped
}

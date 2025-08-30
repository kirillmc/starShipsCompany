package converter

import (
	"github.com/kirillmc/starShipsCompany/order/internal/model"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
)

func SetTransactionUUIDAPI(transactionUUID model.TransactionUUID) orderV1.OptString {
	return orderV1.OptString{
		Value: transactionUUID,
		Set:   true,
	}
}

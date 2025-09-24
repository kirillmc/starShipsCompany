package kafka

import "github.com/kirillmc/starShipsCompany/assembly/internal/model"

type OrderPaidDecoder interface {
	Decode(data []byte) (model.OrderPaidEvent, error)
}

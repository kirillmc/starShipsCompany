package kafka

import "github.com/kirillmc/starShipsCompany/order/internal/model"

type Decoder interface {
	Decode(data []byte) (model.OrderAssembledEvent, error)
}

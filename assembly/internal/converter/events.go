package converter

import "github.com/kirillmc/starShipsCompany/assembly/internal/model"

func ToEventOrderAssembled(event model.OrderPaidEvent, buildTimeSec int64) model.OrderAssembledEvent {
	eventMapped := model.OrderAssembledEvent{
		EventUUID:    event.EventUUID,
		OrderUUID:    event.OrderUUID,
		UserUUID:     event.UserUUID,
		BuildTimeSec: buildTimeSec,
	}

	return eventMapped
}

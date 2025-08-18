package v1

import (
	"context"
	"github.com/google/uuid"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	inventoryV1 "github.com/kirillmc/starShipsCompany/shared/pkg/proto/inventory/v1"
	"net/http"
)

func (a *api) CreateOrder(ctx context.Context, req *orderV1.CreateOrderRequest) (orderV1.CreateOrderRes, error) {
	if req == nil {
		return &orderV1.CreateOrderResponse{}, nil
	}

	inventoryReq := &inventoryV1.ListPartsRequest{Filter: &inventoryV1.PartsFilter{Uuids: req.PartUuids}}
	resp, err := a.inventoryService.ListParts(ctx, inventoryReq)
	if err != nil {
		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}, nil
	}
	if len(resp.Parts) < len(req.PartUuids) {
		return &orderV1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: "internal error",
		}, nil
	}

	var totalPrice float64
	partsUUIDS := make([]string, 0, len(resp.Parts))
	for _, part := range resp.Parts {
		totalPrice += part.Price
		partsUUIDS = append(partsUUIDS, part.Uuid)
	}

	orderUUID := uuid.NewString()

	a.storage.addOrder(orderUUID, req.UserUUID, partsUUIDS, totalPrice)

	return &orderV1.CreateOrderResponse{
		OrderUUID:  orderUUID,
		TotalPrice: totalPrice,
	}, nil
}

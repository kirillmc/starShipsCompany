package converter

import (
	"fmt"

	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
	"github.com/kirillmc/starShipsCompany/order/internal/serviceErrors"
)

func ToModelOrder(ordersWthParts []*repoModel.OrderWthPart) (model.Order, error) {
	if len(ordersWthParts) == 0 {
		return model.Order{}, fmt.Errorf("%w: order not found", serviceErrors.ErrNotFound)
	}

	serviceOrder, err := getNotNilOrder(ordersWthParts)
	if err != nil {
		return model.Order{}, err
	}

	partUUIDS := make([]model.PartUUID, 0, len(ordersWthParts))
	for _, orderWthPart := range ordersWthParts {
		if orderWthPart != nil && orderWthPart.PartUUID != nil {
			partUUIDS = append(partUUIDS, *orderWthPart.PartUUID)
		}
	}
	serviceOrder.PartUUIDs = partUUIDS

	return serviceOrder, nil
}

func getNotNilOrder(ordersWthParts []*repoModel.OrderWthPart) (model.Order, error) {
	for _, orderWthPart := range ordersWthParts {
		if orderWthPart != nil {
			return model.Order{
				OrderUUID:       orderWthPart.OrderUUID,
				UserUUID:        orderWthPart.UserUUID,
				TotalPrice:      orderWthPart.TotalPrice,
				TransactionUUID: orderWthPart.TransactionUUID,
				PaymentMethod:   ToModelPaymentMethod(orderWthPart.PaymentMethod),
				Status:          ToModelOrderStatus(orderWthPart.Status),
			}, nil
		}
	}

	return model.Order{}, fmt.Errorf("%w: order not found", serviceErrors.ErrNotFound)
}

func ToRepoGetOrderParams(params model.GetOrderParams) repoModel.GetOrderParams {
	return repoModel.GetOrderParams{
		OrderUUID: params.OrderUUID,
	}
}

func ToRepoCreateOrder(createOrder model.CreateOrder) repoModel.Order {
	orderMapped := repoModel.Order{
		OrderUUID:  createOrder.OrderUUID,
		UserUUID:   createOrder.UserUUID,
		TotalPrice: createOrder.TotalPrice,
		Status:     repoModel.OrderStatusPendingPayment,
	}

	return orderMapped
}

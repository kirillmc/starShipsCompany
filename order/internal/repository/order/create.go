package order

import (
	"context"

	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/converter"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/model"
)

func (r *repository) Create(_ context.Context, order model.CreateOrder) (model.OrderInfo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderRepo := converter.ToRepoCreateOrder(order)
	r.orders[order.OrderUUID] = &repoModel.Order{
		OrderUUID:       orderRepo.OrderUUID,
		UserUUID:        orderRepo.UserUUID,
		PartUUIDs:       orderRepo.PartUUIDs,
		TotalPrice:      orderRepo.TotalPrice,
		TransactionUUID: orderRepo.TransactionUUID,
		PaymentMethod:   orderRepo.PaymentMethod,
		Status:          orderRepo.Status,
	}

	return converter.ToModelOrderInfo(r.orders[order.OrderUUID].OrderUUID, r.orders[order.OrderUUID].TotalPrice), nil
}

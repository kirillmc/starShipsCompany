package order

import (
	"context"

	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	"github.com/kirillmc/starShipsCompany/order/internal/repository/converter"
)

func (r *repository) UpdateOrder(_ context.Context, updateOrderParams model.UpdateOrderParams) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	updateOrderParamsRepo := converter.ToRepoUpdateOrderParams(updateOrderParams)

	if updateOrderParamsRepo.UserUUID != nil {
		r.orders[updateOrderParamsRepo.OrderUUID].UserUUID = *updateOrderParamsRepo.UserUUID
	}
	if updateOrderParamsRepo.PartUUIDs != nil {
		r.orders[updateOrderParamsRepo.OrderUUID].PartUUIDs = updateOrderParamsRepo.PartUUIDs
	}
	if updateOrderParamsRepo.TotalPrice != nil {
		r.orders[updateOrderParamsRepo.OrderUUID].TotalPrice = *updateOrderParamsRepo.TotalPrice
	}
	if updateOrderParamsRepo.TransactionUUID != nil {
		r.orders[updateOrderParamsRepo.OrderUUID].TransactionUUID = *updateOrderParamsRepo.TransactionUUID
	}
	if updateOrderParamsRepo.PaymentMethod != nil {
		r.orders[updateOrderParamsRepo.OrderUUID].PaymentMethod = *updateOrderParamsRepo.PaymentMethod
	}
	if updateOrderParamsRepo.Status != nil {
		r.orders[updateOrderParamsRepo.OrderUUID].Status = *updateOrderParamsRepo.Status
	}

	return nil
}

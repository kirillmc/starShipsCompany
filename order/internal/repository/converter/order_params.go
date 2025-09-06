package converter

import (
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/model"
	"github.com/samber/lo"
)

func ToRepoUpdateOrderParams(params model.UpdateOrderParams) repoModel.UpdateOrderParams {
	paramsMapped := repoModel.UpdateOrderParams{
		OrderUUID: params.OrderUUID,
	}

	if params.PaymentMethod != nil {
		paramsMapped.PaymentMethod = lo.ToPtr(ToRepoPaymentMethod(*params.PaymentMethod))
	}

	if params.UserUUID != nil {
		paramsMapped.UserUUID = lo.ToPtr(*params.UserUUID)
	}

	if params.PartUUIDs != nil {
		paramsMapped.PartUUIDs = params.PartUUIDs
	}

	if params.TotalPrice != nil {
		paramsMapped.TotalPrice = lo.ToPtr(*params.TotalPrice)
	}

	if params.TransactionUUID != nil {
		paramsMapped.TransactionUUID = lo.ToPtr(*params.TransactionUUID)
	}

	if params.Status != nil {
		paramsMapped.Status = lo.ToPtr(ToRepoOrderStatus(*params.Status))
	}

	return paramsMapped
}

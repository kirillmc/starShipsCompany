package converter

import (
	model "github.com/kirillmc/starShipsCompany/order/internal/model"
	repoModel "github.com/kirillmc/starShipsCompany/order/internal/repository/pg/model"
)

func ToModelOrderStatus(status repoModel.OrderStatus) model.OrderStatus {
	switch status {
	case repoModel.OrderStatusPendingPayment:
		return model.OrderStatusPendingPayment
	case repoModel.OrderStatusPaid:
		return model.OrderStatusPaid
	case repoModel.OrderStatusCancelled:
		return model.OrderStatusCancelled
	default:
		return model.OrderStatusUnspecified
	}
}

func ToRepoOrderStatus(status model.OrderStatus) repoModel.OrderStatus {
	switch status {
	case model.OrderStatusPendingPayment:
		return repoModel.OrderStatusPendingPayment
	case model.OrderStatusPaid:
		return repoModel.OrderStatusPaid
	case model.OrderStatusCancelled:
		return repoModel.OrderStatusCancelled
	default:
		return repoModel.OrderStatusUnspecified
	}
}

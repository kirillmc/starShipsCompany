package v1

import (
	"context"
	orderV1 "github.com/kirillmc/starShipsCompany/shared/pkg/openapi/order/v1"
	"net/http"
)

func (a *api) NewError(_ context.Context, err error) *orderV1.GenericErrorStatusCode {
	return &orderV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: orderV1.GenericError{
			Code:    orderV1.NewOptInt(http.StatusInternalServerError),
			Message: orderV1.NewOptString(err.Error()),
		},
	}
}

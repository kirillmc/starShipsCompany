package serviceErrors

import "errors"

var (
	ErrInternalServer              = errors.New("internal server error")
	ErrNotFound                    = errors.New("not found error")
	ErrNotFoundFromRemoteInventory = errors.New("failed to find parts from inventory")
	ErrOnConflict                  = errors.New("on conflict error")
	ErrUnprocessableEntity         = errors.New("unprocessable entity error")
)

package serviceErrors

import "errors"

var (
	ErrInternalServer      = errors.New("internal server error")
	ErrNotFound            = errors.New("not found error")
	ErrOnConflict          = errors.New("on conflict error")
	ErrUnprocessableEntity = errors.New("unprocessable entity error")
)

package error

import "errors"

var (
	ErrGeneralBadRequest = errors.New("bad request")
	ErrGeneralNotFound   = errors.New("not found")
	ErrGeneralDB         = errors.New("DB error")
)

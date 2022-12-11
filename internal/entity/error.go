package entity

import "errors"

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrNotFound            = errors.New("not found")
	ErrorBadParamInput     = errors.New("provided param is not valid")
)

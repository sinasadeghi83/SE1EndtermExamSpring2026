package account

import "errors"

var (
	ErrNotFound     = errors.New("account service: not found")
	ErrInvalidInput = errors.New("account service: invalid input")
)

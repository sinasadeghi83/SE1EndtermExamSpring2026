package transaction

import "errors"

var (
	ErrNotFound      = errors.New("transaction service: not found")
	ErrInvalidAmount = errors.New("transaction service: invalid amount")
)

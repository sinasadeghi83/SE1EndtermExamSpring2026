package transaction

import "errors"

var (
	ErrNotFound      = errors.New("transaction: not found")
	ErrInvalidAmount = errors.New("transaction: invalid amount")
)

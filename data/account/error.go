package account

import "errors"

var (
	ErrNotFound          = errors.New("account: not found")
	ErrInsufficientFunds = errors.New("account: insufficient funds")
	ErrAccountClosed     = errors.New("account: account is closed")
)

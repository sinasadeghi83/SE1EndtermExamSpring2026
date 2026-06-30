package loan

import "errors"

var (
	ErrNotFound         = errors.New("loan: not found")
	ErrAlreadyEvaluated = errors.New("loan: already evaluated")
	ErrInvalidPrincipal = errors.New("loan: invalid principal amount")
)

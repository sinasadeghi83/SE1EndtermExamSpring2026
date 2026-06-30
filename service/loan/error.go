package loan

import "errors"

var (
	ErrNotFound         = errors.New("loan service: not found")
	ErrAlreadyEvaluated = errors.New("loan service: already evaluated")
)

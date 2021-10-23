package infra

import "errors"

// List of errors
var (
	ErrUnsupportedType = errors.New("infra: unsupported type")
	ErrNotFound        = errors.New("infra: not found")
	ErrDuplicateKey    = errors.New("infra: duplicate key")
	ErrInvalidData     = errors.New("infra: invalid data")
)

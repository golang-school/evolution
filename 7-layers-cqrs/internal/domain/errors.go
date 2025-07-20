package domain

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrUUIDInvalid   = errors.New("invalid UUID format")
)

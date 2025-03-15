package services

import "errors"

var (
	ErrInternal = errors.New("internal server error")
	ErrNotFound = errors.New("not found")
)

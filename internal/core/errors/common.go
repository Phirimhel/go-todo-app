package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argumetn")
	ErrConflict        = errors.New("conflict")
)

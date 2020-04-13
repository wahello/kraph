package errors

import err "errors"

var (
	// ErrNotImplemented is returned when requesting functionality has not been implemented yet
	ErrNotImplemented = err.New("not implemented")
	// ErrUnknownObject is returned when requesting an unknown object
	ErrUnknownObject = err.New("unknown object")
	// ErrUnknownEntity is return when requesting and unknown store entity
	ErrUnknownEntity = err.New("unknown entity")
)
package kraph

import "errors"

var (
	// ErrNotImplemented is returned by functions whose functionality has not been implemented yet
	ErrNotImplemented = errors.New("not implemented")
	// ErrUnknownObject is returned when requesting an object which is not recognised
	ErrUnknownObject = errors.New("unknown object")
)
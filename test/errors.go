package test

import (
	"errors"
	"fmt"
)

var (
	// NoTestInstance is the error returned when no testing instance was
	// provided. Readers must return this error as is and not wrap it as
	// callers are expected to check for this error with ==.
	NoTestInstance error
)

func init() {
	NoTestInstance = errors.New("no testing instance was provided")
}

// ErrPanic is an error that represents a panic.
type ErrPanic struct {
	// Value is the value of the panic.
	Value any
}

// Error implements the error interface.
func (e ErrPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.Value)
}

// NewErrPanic creates a new error that represents a panic.
//
// Parameters:
//   - value: The value of the panic.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"panic: <value>"
//
// where <value> is the value of the panic.
func NewErrPanic(value any) error {
	return &ErrPanic{
		Value: value,
	}
}

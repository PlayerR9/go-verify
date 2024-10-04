package test

import "fmt"

// ErrPanic is an error that represents a panic.
type ErrPanic struct {
	// Value is the value of the panic.
	Value any
}

// Error implements the error interface.
//
// Format:
//
//	"panic: <value>"
//
// where <value> is the value of the panic.
func (e ErrPanic) Error() string {
	return fmt.Sprintf("panic: %v", e.Value)
}

// NewErrPanic creates a new error that represents a panic.
//
// Parameters:
//   - value: The value of the panic.
//
// Returns:
//   - *errors.ErrPanic: The new error. Never returns nil.
func NewErrPanic(value any) *ErrPanic {
	return &ErrPanic{
		Value: value,
	}
}

// IsNil checks if the error is nil.
//
// Returns:
//   - bool: True if the error is nil, false otherwise.
func (e *ErrPanic) IsNil() bool {
	return e == nil
}

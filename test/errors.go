package test

import (
	"errors"
	"fmt"
)

var (
	// ErrNilReceiver occurs when a method is called on a receiver that is nil. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"receiver must not be nil"
	ErrNilReceiver error

	// ErrTestNotImpl occurs when a test is not implemented. This can be checked
	// with the == operator.
	//
	// Format:
	// 	"test not implemented"
	ErrTestNotImpl error
)

func init() {
	ErrTestNotImpl = errors.New("test not implemented")
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrTest occurs when a test failed.
type ErrTest struct {
	// Want is the expected value.
	Want string

	// Got is the actual value.
	Got string
}

// Error implements error.
func (e ErrTest) Error() string {
	return "want " + e.Want + ", got " + e.Got
}

// NewErrTest creates and returns a new ErrTest error with the given expected
// and actual values.
//
// Parameters:
//   - want: The expected value.
//   - got: The actual value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// WARNING: Constructors that uses the FAIL namespaces are recommended to be
// used instead of this. However, if the precise use case is not found in the
// FAIL namespaces, this can be used.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected value.
//   - <got> is the actual value.
func NewErrTest(want, got string) error {
	return &ErrTest{
		Want: want,
		Got:  got,
	}
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

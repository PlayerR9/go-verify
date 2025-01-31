package test

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNilReceiver occurs when a method is called on a receiver that is nil.
	//
	// This error can be checked with the == operator.
	//
	// Format:
	// 	"receiver must not be nil"
	ErrNilReceiver error = errors.New("receiver must not be nil")

	// ErrTestNotImpl occurs when a test is not implemented.
	//
	// // This error can be checked with the == operator.
	//
	// Format:
	// 	"test not implemented"
	ErrTestNotImpl error = errors.New("test not implemented")
)

// ErrTest occurs when a test failed.
type ErrTest struct {
	// Kind is the kind of the value.
	Kind string

	// Want is the expected value.
	Want string

	// Got is the actual value.
	Got string
}

// Error implements error.
func (e ErrTest) Error() string {
	var want, got string

	if e.Want == "" {
		want = "something"
	} else {
		want = e.Want
	}

	if e.Got == "" {
		got = "nothing"
	} else {
		got = e.Got
	}

	if e.Kind == "" {
		return "want " + want + ", got " + got
	}

	var builder strings.Builder

	_, _ = builder.WriteString("want ")
	_, _ = builder.WriteString(e.Kind)
	_, _ = builder.WriteString(" to be ")
	_, _ = builder.WriteString(want)
	_, _ = builder.WriteString(", got ")
	_, _ = builder.WriteString(got)

	str := builder.String()
	return str
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
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected value.
//   - <got> is the actual value.
//
// If the kind is empty, "want <want>, got <got>" is used.
func NewErrTest(kind, want, got string) error {
	err := &ErrTest{
		Kind: kind,
		Want: want,
		Got:  got,
	}

	return err
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
	err := &ErrPanic{
		Value: value,
	}

	return err
}

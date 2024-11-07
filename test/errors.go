package test

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// NoTestInstance occurs when a nil testing instance was provided.
	//
	// Format:
	//   "no testing instance was provided"
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

type ErrTestFailed struct {
	Want string
	Got  error
}

func (e ErrTestFailed) Error() string {
	var got string

	if e.Got == nil {
		got = "no error"
	} else {
		got = e.Got.Error()
	}

	var msg string

	if e.Want == "" {
		msg = "want no error, got "
	} else {
		msg = "want " + strconv.Quote(e.Want) + ", got "
	}

	return msg + strconv.Quote(got)
}

func NewErrTestFailed() error {
	return &ErrTestFailed{}
}

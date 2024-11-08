package common

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var (
	// ErrInvalidBase occurs when the base is zero.
	//
	// Format:
	//   "base must not be zero"
	ErrInvalidBase error

	// ErrNilReceiver occurs when the receiver is nil.
	//
	// Format:
	//   "receiver must not be nil"
	ErrNilReceiver error

	// NoTestInstance occurs when a nil testing instance was provided.
	//
	// Format:
	//   "no testing instance was provided"
	NoTestInstance error
)

func init() {
	ErrInvalidBase = errors.New("base must not be zero")
	ErrNilReceiver = errors.New("receiver must not be nil")
	NoTestInstance = errors.New("a testing instance was be provided")
}

// ErrTestFailed is an error that represents a test failure.
type ErrTestFailed struct {
	// Want is the expected value.
	Want string

	// Got is the actual value encountered.
	Got string
}

// Error implements the error interface.
func (e ErrTestFailed) Error() string {
	return "want " + e.Want + ", got " + e.Got
}

// NewErrTestFailed creates a new error that represents a test failure. This should only be used
// when there are no FAIL.Wrong* functions that suit your use case.
//
// Parameters:
//   - want: The expected value.
//   - got: The actual value encountered.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// where:
//   - <want> is the expected value.
//   - <got> is the actual value.
func NewErrTestFailed(want string, got string) error {
	return &ErrTestFailed{
		Want: want,
		Got:  got,
	}
}

// failT is for private use only.
type failT struct{}

// FAIL is the namespace for making errors.
var FAIL failT

func init() {
	FAIL = failT{}
}

// WrongError reports a test failure when the actual error does not match the
// expected error string. It logs an error message to the provided testing instance.
//
// Parameters:
//   - t: The testing instance used to report the error. Must not be nil.
//   - want: The expected error message.
//   - got: The actual error encountered.
//
// Panics:
//   - NoTestInstance: If the testing instance `t` is nil.
//
// Logs:
//   - An error message in the format "want <want>, got <got>", where <want> is
//     the expected error message (quoted if not empty), and <got> is the actual error message
//     (quoted if not nil), otherwise "no error".
func (failT) WrongError(t *testing.T, want string, got error) {
	if t == nil {
		panic(NoTestInstance)
	}

	var got_str string

	if got == nil {
		got_str = "no error"
	} else {
		got_str = strconv.Quote(got.Error())
	}

	t.Error(&ErrTestFailed{
		Want: OrQuoteElse(want, "no error"),
		Got:  got_str,
	})
}

// WrongInt reports a test failure when the actual integer does not match the
// expected integer. It logs an error message to the provided testing instance.
//
// Parameters:
//   - t: The testing instance used to report the error. Must not be nil.
//   - want: The expected integer.
//   - got: The actual integer encountered.
//
// Panics:
//   - NoTestInstance: If the testing instance `t` is nil.
//
// Logs:
//   - An error message in the format "want <want>, got <got>", where <want> is
//     the expected integer, and <got> is the actual integer.
func (failT) WrongInt(t *testing.T, want int, got int) {
	if t == nil {
		panic(NoTestInstance)
	}

	t.Error(&ErrTestFailed{
		Want: strconv.Itoa(want),
		Got:  strconv.Itoa(got),
	})
}

// WrongBool reports a test failure when the actual boolean does not match the
// expected boolean. It logs an error message to the provided testing instance.
//
// Parameters:
//   - t: The testing instance used to report the error. Must not be nil.
//   - want: The expected boolean.
//   - got: The actual boolean encountered.
//
// Panics:
//   - NoTestInstance: If the testing instance `t` is nil.
//
// Logs:
//   - An error message in the format "want <want>, got <got>", where <want> is
//     the expected boolean, and <got> is the actual boolean.
func (failT) WrongBool(t *testing.T, want bool, got bool) {
	if t == nil {
		panic(NoTestInstance)
	}

	t.Error(&ErrTestFailed{
		Want: strconv.FormatBool(want),
		Got:  strconv.FormatBool(got),
	})
}

// WrongAny reports a test failure when the actual value does not match the
// expected value. It logs an error message to the provided testing instance.
//
// Parameters:
//   - t: The testing instance used to report the error. Must not be nil.
//   - want: The expected value.
//   - got: The actual value encountered.
//
// Panics:
//   - NoTestInstance: If the testing instance `t` is nil.
//
// Logs:
//   - An error message in the format "want <want>, got <got>", where <want> is
//     the expected value, and <got> is the actual value.
func (failT) WrongAny(t *testing.T, want any, got any) {
	if t == nil {
		panic(NoTestInstance)
	}

	t.Error(&ErrTestFailed{
		Want: fmt.Sprint(want),
		Got:  fmt.Sprint(got),
	})
}

// CheckErr compares the expected error message with the actual error encountered
// during testing. It reports a test failure if the actual error does not match
// the expected error message.
//
// Parameters:
//   - t: The testing instance used to report the error. Must not be nil.
//   - expected: The expected error message.
//   - got: The actual error encountered.
//
// Returns:
//   - bool: True if the actual error matches the expected error message, false otherwise.
//
// Panics:
//   - NoTestInstance: If the testing instance `t` is nil.
//
// Logs:
//   - An error message in the format "want <expected>, got <got>", where <expected>
//     is the expected error message (quoted if not empty), and <got> is the actual
//     error message (quoted if not nil), otherwise "no error".
func (failT) CheckErr(t *testing.T, expected string, got error) bool {
	if t == nil {
		panic(NoTestInstance)
	}

	if expected != "" {
		if got == nil || got.Error() != expected {
			FAIL.WrongError(t, expected, got)
			return false
		}

		return true
	} else if got == nil {
		return true
	} else {
		FAIL.WrongError(t, "", got)
		return false
	}
}

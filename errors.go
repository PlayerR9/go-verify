package errs

import (
	"strings"

	"github.com/PlayerR9/go-verify/errkit"
)

// ErrorCode is the type of an error code.
type ErrorCode int

const (
	// InvalidOperation occurs when an operation (such as a function call) fails
	// or cannot be performed for a reason or another.
	InvalidOperation ErrorCode = iota

	// UnexpectedValue occurs when a value (or type) was expected but, instead,
	// another one was encountered.
	UnexpectedValue
)

// NewNilReceiver creates a new errors.Err error with the code InvalidOperation and the
// message "receiver must not be nil".
//
// Returns:
//   - *errors.Err[ErrorCode]: The new error. Never returns nil.
func NewNilReceiver() *errkit.CodedErr[ErrorCode] {
	err := errkit.New(InvalidOperation, "receiver must not be nil")

	return err
}

// NewUnsupportedValue creates a new errors.Err error with the code InvalidOperation and the
// message "value of <expected> is not a supported <kind> type".
//
// Parameters:
//   - kind: The kind of the value. Ignored if not provided.
//   - expected: The expected value. Ignored if not provided.
//
// Returns:
//   - *errors.Err[ErrorCode]: The new error. Never returns nil.
func NewUnsupportedValue(kind, expected string) *errkit.CodedErr[ErrorCode] {
	var builder strings.Builder

	builder.WriteString("value ")

	if expected != "" {
		builder.WriteString("of ")
		builder.WriteString(expected)
	}

	builder.WriteString(" is not  ")

	if kind != "" {
		builder.WriteString("a supported ")
		builder.WriteString(kind)
		builder.WriteString(" type")
	} else {
		builder.WriteString("supported")
	}

	err := errkit.New(InvalidOperation, builder.String())

	return err
}

// NewUnexpectedType creates a new errors.Err error with the code UnexpectedValue and the
// message "expect <kind> to be of type <expected>, got <got> instead".
//
// Parameters:
//   - kind: The kind of the value. If not provided, "something" is used.
//   - expected: The expected type of the value.
//   - got: The actual type of the value.
func NewUnexpectedType(kind string, expected, got any) *errkit.CodedErr[ErrorCode] {
	var builder strings.Builder

	builder.WriteString("expected ")
	OrElse(&builder, kind, "something")
	builder.WriteString("to be ")

	if expected == nil {
		builder.WriteString("nil")
	} else {
		builder.WriteString("of type ")
		StringOfType(&builder, expected)
	}

	builder.WriteString(", got ")

	if got == nil {
		builder.WriteString("no type")
	} else {
		StringOfType(&builder, got)
	}

	builder.WriteString(" instead")

	err := errkit.New(UnexpectedValue, builder.String())

	return err
}

// NewNotAsExpected creates a new errors.Err error with the code UnexpectedValue and the
// message "expected <kind> to be <expected>, got <got> instead".
//
// Parameters:
//   - kind: The kind of the value. Ignored if not provided.
//   - expected: The expected value. If not provided, "nothing" is used.
//   - got: The actual value. If not provided, "nothing" is used.
//
// Returns:
//   - *errors.Err[ErrorCode]: The new error. Never returns nil.
func NewNotAsExpected(kind, expected, got string) *errkit.CodedErr[ErrorCode] {
	var builder strings.Builder

	builder.WriteString("expected ")

	if kind != "" {
		builder.WriteString(kind)
		builder.WriteString(" to be ")
	}

	OrElse(&builder, expected, "nothing")
	builder.WriteString(", got ")
	OrElse(&builder, got, "nothing")
	builder.WriteString(" instead")

	err := errkit.New(UnexpectedValue, builder.String())

	return err
}

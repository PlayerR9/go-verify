package errs

import (
	"strconv"
	"strings"

	"github.com/PlayerR9/go-verify/errkit"
)

// ErrorCode is the type of an error code.
type ErrorCode int

const (
	// BadParameter occurs when a parameter is invalid or is not
	// valid for some reason. For example, a nil pointer when nil
	// pointers are not allowed.
	BadParameter ErrorCode = iota

	// InvalidUsage occurs when users call a function without
	// proper setups or preconditions.
	InvalidUsage

	// NoSuchKey occurs when a context key is requested but does
	// not exist.
	NoSuchKey

	// OperationFail occurs when an operation cannot be completed
	// due to an internal error.
	OperationFail

	// UnexpectedValue occurs when a value (or type) was expected but, instead,
	// another one was encountered.
	UnexpectedValue
)

// NewErrNilReceiver creates a new error.Err error with the code
// OperationFail.
//
// Parameters:
//   - frame: The frame of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilReceiver(frame string) *errkit.CodedErr[ErrorCode] {
	err := errkit.New(OperationFail, "receiver must not be nil")
	err.AddSuggestion("Did you forget to initialize the receiver?")

	err.AddFrame(frame)

	return err
}

// NewErrInvalidParameter creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - message: The message of the error.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
//
// This function is mostly useless since it just wraps BadParameter.
func NewErrInvalidParameter(frame, message string) *errkit.CodedErr[ErrorCode] {
	err := errkit.New(BadParameter, message)
	err.AddFrame(frame)

	return err
}

// NewErrNilParameter creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - parameter: the name of the invalid parameter.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNilParameter(frame, parameter string) *errkit.CodedErr[ErrorCode] {
	msg := "parameter (" + strconv.Quote(parameter) + ") must not be nil"

	err := errkit.New(BadParameter, msg)
	err.AddSuggestion("Maybe you forgot to initialize the parameter?")

	return err
}

// NewErrInvalidUsage creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - message: The message of the error.
//   - usage: The usage/suggestion to solve the problem.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrInvalidUsage(frame, message, usage string) *errkit.CodedErr[ErrorCode] {
	err := errkit.New(InvalidUsage, message)

	err.AddSuggestion(usage)

	err.AddFrame(frame)

	return err
}

// NewErrNoSuchKey creates a new error.Err error.
//
// Parameters:
//   - frame: The frame of the error.
//   - key: The key that does not exist.
//
// Returns:
//   - *error.Err: The new error. Never returns nil.
func NewErrNoSuchKey(frame, key string) *errkit.CodedErr[ErrorCode] {
	err := errkit.New(NoSuchKey, "key ("+strconv.Quote(key)+") does not exist")

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

	err := errkit.New(OperationFail, builder.String())

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

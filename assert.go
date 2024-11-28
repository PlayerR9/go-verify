package assert

import (
	"fmt"
)

// Cond checks a condition and if it is not true, panics with ErrAssertFail.
//
// Parameters:
//   - cond: The condition to check.
//   - msg: The error message to use if the condition is not true.
//
// Panics:
//   - ErrAssertFail: If the condition is not true.
func Cond(cond bool, msg string) {
	if cond {
		return
	}

	err := NewErrAssertFail(msg)
	panic(err)
}

// Condf checks a condition and if it is not true, panics with ErrAssertFail.
//
// Parameters:
//   - cond: The condition to check.
//   - format: The format string for the error message to use if the condition is not true.
//   - args: The arguments to use with the format string.
//
// Panics:
//   - ErrAssertFail: If the condition is not true.
func Condf(cond bool, format string, args ...any) {
	if cond {
		return
	}

	err := NewErrAssertFail(fmt.Sprintf(format, args...))
	panic(err)
}

// Err panics with ErrAssertFail if the given inner error is not nil. The
// error message will be the given format string with the given arguments
// followed by " = " and the error message of the inner error. If the given
// format string is empty, the error message will be "func() = " followed by the
// error message of the inner error.
//
// Parameters:
//   - inner: The inner error to check.
//   - format: The format string for the error message.
//   - args: The arguments to use with the format string.
//
// Panics:
//   - ErrAssertFail: If the inner error is not nil.
func Err(inner error, format string, args ...any) {
	if inner == nil {
		return
	}

	msg := fmt.Sprintf(format, args...)
	if msg == "" {
		msg = "func()"
	}

	err := NewErrAssertFail(msg + " = " + inner.Error())
	panic(err)
}

/* func X(inner error, name string) {
	if inner == nil {
		return
	}

	var msg string

	if name == "" {
		msg = ""
	} else {
		msg = name
	}

	err := NewErrAssertFail(msg + " = " + inner.Error())
	panic(err)
} */

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

// CheckPanic executes the given function and checks that, not only does it panic, but
// also that the paniced error's message matches the given expected error string.
//
// Parameters:
//   - expected: The expected error string.
//   - fn: The function to execute.
//
// Returns:
//   - error: an error if the fn is nil, it did not panic, or the paniced error's
//     message does not match the expected error string, nil otherwise.
//
// Example:
//
//	err := CheckPanic("something went wrong", func() {
//		panic(errors.New("something went wrong"))
//	})
//
//	if err != nil {
//		t.Error(err) // Does not error.
//	}
func CheckPanic(expected string, fn func()) error {
	if fn == nil {
		return errors.New("no function was provided")
	}

	var err error

	try(&err, fn)

	if err == nil {
		return fmt.Errorf("want %q, got nil", expected)
	}

	got := err.Error()

	if got != expected {
		return fmt.Errorf("want %q, got %q", expected, got)
	}

	return nil
}

// CheckErr checks that the given error matches the given expected message. If not, an error is returned.
//
// Parameters:
//   - expected: The expected error message.
//   - got: The error to check.
//
// Returns:
//   - error: an error if the error does not match the expected message, nil otherwise.
//
// Example:
//
//	err := CheckErr("something went wrong", errors.New("something went wrong"))
//	if err != nil {
//		t.Error(err) // Does not error.
//	}
func CheckErr(expected string, got error) error {
	if got == nil {
		return fmt.Errorf("want %q, got nil", expected)
	}

	msg := got.Error()

	if msg != expected {
		return fmt.Errorf("want %q, got %q", expected, msg)
	}

	return nil
}

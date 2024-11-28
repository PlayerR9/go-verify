package test

import (
	"errors"
)

// try executes the given function and captures any panic that occurs,
// converting it into an error and storing it in the provided error pointer.
//
// Parameters:
//   - err: A pointer to an error variable where the captured panic error
//     will be stored. If the function does not panic, the error remains unchanged.
//   - fn: The function to be executed, which may or may not panic.
func try(err *error, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case string:
			*err = errors.New(r)
		case error:
			*err = r
		default:
			*err = NewErrPanic(r)
		}
	}()

	fn()
}

// Try executes the given function and captures any panic that occurs,
// converting it into an error and returning it.
//
// Parameters:
//   - fn: The function to be executed, which may or may not panic.
//
// Returns:
//   - error: The captured panic error, if any. Otherwise, nil.
//
// If the panic value is a string, it is converted to an error. If the
// panic value is an error, it is returned as-is. In any other case, a new
// ErrPanic is created.
func Try(fn func()) error {
	if fn == nil {
		return nil
	}

	var err error

	try(&err, fn)

	return err
}

// EqualsErr checks if two errors are equal.
//
// Two errors are considered equal iff they are both non-nil and satisfies
// at least one of the following conditions:
//   - One of them is equal to the other according to the error.Is function.
//   - They have the same error message.
//
// Parameters:
//   - err1: The first error.
//   - err2: The second error.
//
// Returns:
//   - bool: True if the errors are equal, false otherwise.
//
// Example:
//
//	err1 := errors.New("this is an error")
//
//	EqualsErr(err1, err1) // true
func EqualsErr(err1, err2 error) bool {
	if err1 == nil || err2 == nil {
		return false
	}

	if errors.Is(err1, err2) || errors.Is(err2, err1) {
		return true
	}

	return err1.Error() == err2.Error()
}

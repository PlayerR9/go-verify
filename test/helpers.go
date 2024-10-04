package test

import (
	"errors"
)

// try executes the given function and stores any paniced error in err.
//
// Parameters:
//   - err: The error to store the paniced error in. If the given function
//     does not panic, the value of err is not changed.
//   - fn: The function to execute.
func try(err *error, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		v, ok := r.(error)
		if !ok {
			(*err) = NewErrPanic(r)
		} else {
			(*err) = v
		}
	}()

	fn()
}

// Try executes the given function and returns any paniced error. If the given
// function does not panic, nil is returned.
//
// Parameters:
//   - fn: The function to execute. If nil, nil is returned.
//
// Returns:
//   - error: The paniced error, if any. Otherwise, nil.
//
// Errors:
//   - *ErrPanic: If the panic value is not an error.
//   - any other error: If the panic value is an error.
//
// Example:
//
//	err := Try(func() {
//		panic("something went wrong")
//	})
//
//	fmt.Println(err) // Prints: something went wrong
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
	return (err1 != nil && err2 != nil) && (errors.Is(err1, err2) ||
		errors.Is(err2, err1) ||
		err1.Error() == err2.Error())
}

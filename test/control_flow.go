package test

import "errors"

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
// If the panic value is an error, it is returned as-is. If the panic value is a string, it is
// converted to an error. In any other case, a new ErrPanic is created.
//
// Example:
//
//	err := Try(func() {
//		panic("something went wrong")
//	})
//
//	fmt.Println(err) // Prints: "something went wrong"
func Try(fn func()) error {
	if fn == nil {
		return nil
	}

	var err error

	try(&err, fn)

	return err
}

package errs

import (
	"github.com/PlayerR9/go-verify/errkit"
)

// AsWithCode returns the error if it is of type C.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - *Err: The error if it is of type T, nil otherwise.
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).

func AsWithCode[C errkit.ErrorCoder](err error, code C) (*errkit.CodedErr[C], bool) {
	if err == nil {
		return nil, false
	}

	stack := []error{err}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		e, ok := top.(*errkit.CodedErr[C])
		if ok && e.Code == code {
			return e, true
		}

		switch top := top.(type) {
		case interface{ Unwrap() error }:
			err := top.Unwrap()
			if err != nil {
				stack = append(stack, err)
			}
		case interface{ Unwrap() []error }:
			for _, err := range top.Unwrap() {
				if err != nil {
					stack = append(stack, err)
				}
			}
		}
	}

	return nil, false
}

// Value is a function that returns the value of the context with the given key.
//
// Parameters:
//   - err: The error to get the value from.
//   - key: The key of the context.
//
// Returns:
//   - T: The value of the context with the given key.
//   - bool: True if the key exists, false otherwise.
func Value[T any](err error, key string) (T, bool) {
	zero := *new(T)

	if err == nil {
		return zero, false
	}

	e, ok := err.(interface {
		Get(key string) (any, bool)
	})
	if !ok {
		return zero, false
	}

	val, ok := e.Get(key)
	if !ok {
		return zero, false
	}

	res, ok := val.(T)
	return res, ok
}

/*
// Is is function that checks if an error is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func Is[T ErrorCoder](err error, code T) bool {
	if err == nil {
		return false
	}

	var sub_err *Err

	ok := errors.As(err, &sub_err)
	if !ok {
		return false
	}

	other, ok := sub_err.Code.(T)
	return ok && other.Int() == code.Int()
} */

/* // As returns the error if it is of type T.
//
// Parameters:
//   - err: The error to check.
//   - code: The error code to check.
//
// Returns:
//   - *Err: The error if it is of type T, nil otherwise.
//   - bool: true if the error is of type T, false otherwise (including if the error is nil).
func As(err error) (*Err, bool) {
	if err == nil {
		return nil, false
	}

	var sub_err *Err

	ok := errors.As(err, &sub_err)
	if !ok {
		return nil, false
	}

	return sub_err, true
} */

/*
// LimitErrorMsg is a function that limits the number of errors in an error chain.
//
// Parameters:
//   - err: The error to limit.
//   - limit: The maximum number of errors to limit.
//
// Returns:
//   - error: The limited error.
//
// If the error is nil or the limit is less than or equal to 0, the function returns nil.
func LimitErrorMsg(err error, limit int) error {
	if err == nil || limit <= 0 {
		return nil
	}

	target := err

	for i := 0; i < limit; i++ {
		w, ok := target.(Unwrapper)
		if !ok {
			return err
		}

		reason := w.Unwrap()
		if reason == nil {
			return err
		}

		target = reason
	}

	if target == nil {
		return err
	}

	w, ok := target.(Unwrapper)
	if !ok {
		return err
	}

	w.ChangeReason(nil)

	return err
} */

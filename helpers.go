package errs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/PlayerR9/go-verify/assert"
	"github.com/PlayerR9/go-verify/errkit"
)

// OrElse is a helper function that either writes the first string
// to the builder if it is not empty, or the second string if it
// is empty.
//
// Parameters:
//   - b: The builder to write to.
//   - first: The first string to write.
//   - second: The second string to write.
//
// Does nothing if the builder is nil.
func OrElse(b *strings.Builder, first, second string) {
	if b == nil {
		return
	}

	if first == "" {
		b.WriteString(second)
	} else {
		b.WriteString(first)
	}
}

// StringOfType is a helper function that either writes the string
// representation of the value to the builder if it is not nil, or
// the string "nil" if it is nil.
//
// Parameters:
//   - b: The builder to write to.
//   - v: The value to write.
//
// Does nothing if the builder is nil.
func StringOfType(b *strings.Builder, v any) {
	if b == nil {
		return
	}

	if v == nil {
		b.WriteString("nil")
	}

	type_of := reflect.TypeOf(v)
	assert.Cond(type_of != nil, "type_of = nil")

	b.WriteString(type_of.String())
}

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

	var sub_err *errkit.CodedErr[C]

	ok := errors.As(err, &sub_err)
	if !ok {
		return nil, false
	}

	if sub_err.Code == code {
		return sub_err, true
	} else {
		return nil, false
	}
}

// Value is a function that returns the value of the context with the given key.
//
// Parameters:
//   - e: The error to get the value from.
//   - key: The key of the context.
//
// Returns:
//   - T: The value of the context with the given key.
//   - error: The error that occurred while getting the value.
func Value[C errkit.ErrorCoder, T any](e *errkit.CodedErr[C], key string) (T, error) {
	zero := *new(T)

	if e == nil || len(e.Context) == 0 {
		return zero, NewErrNoSuchKey("Value()", key)
	}

	x, ok := e.Context[key]
	if !ok {
		return zero, NewErrNoSuchKey("Value()", key)
	}

	if x == nil {
		err := NewErrNoSuchKey("Value()", key)
		err.AddSuggestion("Found a key with the same name but has a nil value")

		return zero, err
	}

	val, ok := x.(T)
	if !ok {
		err := NewErrNoSuchKey("Value()", key)
		err.AddSuggestion(fmt.Sprintf("Found a key with the same name but has a value of type %T", x))

		return zero, err
	}

	return val, nil
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

package assert

import (
	"fmt"
)

// MustCall asserts whether the function does not return an error for the given argument.
// If the function returns an error, it panics with an ErrAssertFailed error.
//
// Parameters:
//   - arg: The argument to pass to the function.
//   - fn: The function to execute.
//   - format: The format of the function call that returned the error.
//   - args: The arguments of the function call.
func MustCall[T any](arg T, fn func(arg T) error, format string, args ...any) {
	if fn == nil {
		panic(NewErrAssertFailed("no function provided"))
	}

	err := fn(arg)
	if err == nil {
		return
	}

	msg := fmt.Sprintf(format, args...)
	msg += " = " + err.Error()

	panic(NewErrAssertFailed(msg))
}

// Must is a helper function that wraps a call to a function that returns (T, error) and
// panics if the error is not nil.
//
// This function is intended to be used to handle errors in a way that is easy to read and write.
//
// Parameters:
//   - res: The result of the function.
//   - err: The error returned by the function.
//
// Returns:
//   - T: The result of the function.
func Must[T comparable](res T, err error) T {
	if err != nil {
		panic(NewErrAssertFailed("err = " + err.Error()))
	} else if res == *new(T) {
		panic(NewErrAssertFailed("res = nil"))
	}

	return res
}

// New is a syntactic sugar asserting constructors. It asserts whether the
// constructor does not return an error and the result is non-nil. If not, it
// panics with an ErrAssertFailed error.
//
// Parameters:
//   - res: The result of the constructor.
//   - inner: The error returned by the constructor.
//
// Example:
//
//	type MyStruct struct {}
//
//	func (my_struct *MyStruct) IsNil() bool {
//		return my_struct == nil
//	}
//
//	func NewMyStruct() (*MyStruct, error) {
//		return nil, nil
//	}
//
//	res := New(NewMyStruct()) // Panics: *MyStruct = nil
func New[T interface{ IsNil() bool }](res T, err error) T {
	if err != nil {
		panic(NewErrAssertFailed("err = " + err.Error()))
	}

	if res.IsNil() {
		msg := fmt.Sprintf("%T = nil", *new(T))
		panic(NewErrAssertFailed(msg))
	}

	return res
}

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

// True checks whether a boolean condition is true. If not, it panics with ErrAssertFail
// using the given format string and arguments. If the format string is empty,
// "func() = false" is used as the error message.
//
// Parameters:
//   - ok: The boolean condition to check.
//   - format: The format string for the error message if the condition is false.
//   - args: The arguments to use with the format string.
//
// Panics:
//   - ErrAssertFail: If the condition is false.
func True(ok bool, format string, args ...any) {
	if ok {
		return
	}

	msg := fmt.Sprintf(format, args...)
	if msg == "" {
		msg = "func()"
	}

	err := NewErrAssertFail(msg + " = false")
	panic(err)
}

// False checks whether a boolean condition is false. If not, it panics with ErrAssertFail
// using the given format string and arguments. If the format string is empty,
// "func() = true" is used as the error message.
//
// Parameters:
//   - ok: The boolean condition to check.
//   - format: The format string for the error message if the condition is true.
//   - args: The arguments to use with the format string.
//
// Panics:
//   - ErrAssertFail: If the condition is true.
func False(ok bool, format string, args ...any) {
	if !ok {
		return
	}

	msg := fmt.Sprintf(format, args...)
	if msg == "" {
		msg = "func()"
	}

	err := NewErrAssertFail(msg + " = true")
	panic(err)
}

// NotZero asserts whether the variable is not its zero value. If not, it
// panics with an ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, "variable" is used.
//
// Example:
//
//	v := 0
//
//	NotZero[int](v, "v") // Panics: v is zero
func NotZero[E comparable](v E, name string) {
	if v != *new(E) {
		return
	}

	if name == "" {
		name = "variable"
	}

	err := NewErrAssertFail(name + " is zero")
	panic(err)
}

// Type asserts whether the variable is of type E. If not, it panics with an
// ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//   - allow_nil: Whether to allow the variable to be nil.
//
// Example:
//
//	v := "foo"
//	Type[int](v, "v", false) // Panics: v = string, want int
func Type[E any](v any, name string, allow_nil bool) {
	if name == "" {
		name = "variable"
	}

	if v == nil && !allow_nil {
		err := NewErrAssertFail(name + " = nil")
		panic(err)
	} else if v == nil {
		return
	}

	_, ok := v.(E)
	if ok {
		return
	}

	msg := fmt.Sprintf("%s = %T, want %T", name, v, *new(E))

	err := NewErrAssertFail(msg)
	panic(err)
}

// Deref asserts whether the variable is both non-nil and is of type E. If
// not, it panics with an ErrAssertFailed error.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//
// Returns:
//   - E: The dereferenced variable.
//
// Example:
//
//	var v *int
//	_ = Deref[string](v, "v") // Panics: v = *int, want string
func Deref[E any](v *E, name string) E {
	if name == "" {
		name = "variable"
	}

	if v != nil {
		return *v
	}

	msg := fmt.Sprintf("%s = nil, want %T", name, *new(E))

	err := NewErrAssertFail(msg)
	panic(err)
}

// Conv asserts whether the variable is of type E. If not, it panics with an
// ErrAssertFailed error. Unlike with Type(), this returns the converted
// type as well.
//
// Parameters:
//   - v: The variable to assert.
//   - name: The name of the variable. If empty, the name "variable" is used.
//
// Example:
//
//	v := "foo"
//	res := Conv[int](v, "v") // Panics: v = string, want int
func Conv[E any](v any, name string) E {
	if name == "" {
		name = "variable"
	}

	if v == nil {
		err := NewErrAssertFail(name + " = nil")
		panic(err)
	}

	val, ok := v.(E)
	if ok {
		return val
	}

	msg := fmt.Sprintf("%s = %T, want %T", name, v, *new(E))

	err := NewErrAssertFail(msg)
	panic(err)
}

// NotNil asserts whether the variable is not nil. If it is nil, it panics
// with an ErrAssertFail error.
//
// Parameters:
//   - v: The pointer to the variable to assert.
//   - name: The name of the variable. If empty, "variable" is used.
//
// Panics:
//   - ErrAssertFail: If the variable is nil.
//
// Example:
//
//	var v *int
//	NotNil[int](v, "v") // Panics: v = nil
func NotNil[E any](v *E, name string) {
	if v != nil {
		return
	}

	err := NewErrAssertFail(name + " = nil")
	panic(err)
}

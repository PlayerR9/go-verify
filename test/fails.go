package test

import (
	"errors"
	"fmt"
	"strconv"
)

// failT is for private use only
type failT struct{}

var (
	// FAIL is the namespace for creating ErrTest errors.
	FAIL failT
)

func init() {
	FAIL = failT{}
}

// String creates and returns a new ErrTest error with the given expected and
// actual string values. The string values are quoted.
//
// Parameters:
//   - want: The expected string value.
//   - got: The actual string value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected quoted string value.
//   - <got> is the actual quoted string value.
func (failT) String(want, got string) error {
	return &ErrTest{
		Want: strconv.Quote(want),
		Got:  strconv.Quote(got),
	}
}

// Int creates and returns a new ErrTest error with the given expected and
// actual integer values.
//
// Parameters:
//   - want: The expected integer value.
//   - got: The actual integer value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected integer value.
//   - <got> is the actual integer value.
func (failT) Int(want, got int) error {
	return &ErrTest{
		Want: strconv.Itoa(want),
		Got:  strconv.Itoa(got),
	}
}

// Uint creates and returns a new ErrTest error with the given expected and
// actual integer values.
//
// Parameters:
//   - want: The expected integer value.
//   - got: The actual integer value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected unsigned integer value.
//   - <got> is the actual unsigned integer value.
func (failT) Uint(want, got uint) error {
	return &ErrTest{
		Want: strconv.FormatUint(uint64(want), 10),
		Got:  strconv.FormatUint(uint64(got), 10),
	}
}

// Err creates and returns a new ErrTest error with the given expected and
// actual values. Error messages are quoted.
//
// Parameters:
//   - want: The expected error.
//   - got: The actual error.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected error message if want is not nil, "no error"
//     otherwise.
//   - <got> is the actual error message if got is not nil, "no error"
//     otherwise.
func (failT) Err(want, got error) error {
	if want == nil && got == nil {
		return nil
	}

	if want == nil {
		return &ErrTest{
			Want: "no error",
			Got:  strconv.Quote(got.Error()),
		}
	}

	if got == nil {
		return &ErrTest{
			Want: strconv.Quote(want.Error()),
			Got:  "no error",
		}
	}

	ok := errors.Is(want, got)
	if ok {
		return nil
	}

	ok = errors.Is(got, want)
	if ok {
		return nil
	}

	want_str := want.Error()
	got_str := got.Error()

	if want_str == got_str {
		return nil
	}

	return &ErrTest{
		Want: strconv.Quote(want_str),
		Got:  strconv.Quote(got_str),
	}
}

// Any creates and returns a new ErrTest error with the given expected and
// actual values.
//
// Parameters:
//   - want: The expected value.
//   - got: The actual value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected value if want is not nil, "nothing" otherwise.
//   - <got> is the actual value if got is not nil, "nothing" otherwise.
func (failT) Any(want, got any) error {
	var want_str, got_str string

	if want == nil {
		want_str = "nothing"
	} else {
		want_str = fmt.Sprint(want)
	}

	if got == nil {
		got_str = "nothing"
	} else {
		got_str = fmt.Sprint(got)
	}

	return &ErrTest{
		Want: want_str,
		Got:  got_str,
	}
}

// ErrorMessage creates and returns a new ErrTest error with the given expected and
// actual error messages.
//
// Parameters:
//   - got: The actual error encountered.
//   - want: The expected error message.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Returns nil if the actual
//     error matches the expected error message, or if both the actual and expected
//     errors are nil.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected error message if want is not empty, "no error"
//     otherwise.
//   - <got> is the actual error message if got is not nil, "no error" otherwise.
func (failT) ErrorMessage(got error, want string) error {
	if got == nil {
		if want == "" {
			return nil
		}

		return &ErrTest{
			Want: strconv.Quote(want),
			Got:  "no error",
		}
	}

	got_str := got.Error()
	if want == got_str {
		return nil
	}

	var want_str string

	if want != "" {
		want_str = strconv.Quote(want)
	} else {
		want_str = "no error"
	}

	return &ErrTest{
		Want: want_str,
		Got:  strconv.Quote(got_str),
	}
}

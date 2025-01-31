package test

import (
	"fmt"
	"strconv"
)

const (
	// Base10 is the base 10 integer.
	Base10 int = 10
)

// failT is for private use only
type failT struct{}

var (
	// FAIL is the namespace for creating ErrTest errors.
	FAIL failT = failT{}
)

// String creates and returns a new ErrTest error with the given expected and
// actual string values. The string values are quoted.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected string value.
//   - got: The actual string value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected quoted string value.
//   - <got> is the actual quoted string value.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) String(kind, want, got string) error {
	if want != "" {
		want = strconv.Quote(want)
	}

	if got != "" {
		got = strconv.Quote(got)
	}

	err := &ErrTest{
		Kind: kind,
		Want: want,
		Got:  got,
	}

	return err
}

// Int creates and returns a new ErrTest error with the given expected and
// actual integer values.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected integer value.
//   - got: The actual integer value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected integer value.
//   - <got> is the actual integer value.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) Int(kind string, want, got int) error {
	want_str := strconv.Itoa(want)
	got_str := strconv.Itoa(got)

	err := &ErrTest{
		Kind: kind,
		Want: want_str,
		Got:  got_str,
	}

	return err
}

// Uint creates and returns a new ErrTest error with the given expected and
// actual integer values.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected integer value.
//   - got: The actual integer value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected unsigned integer value.
//   - <got> is the actual unsigned integer value.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) Uint(kind string, want, got uint) error {
	want_str := strconv.FormatUint(uint64(want), Base10)
	got_str := strconv.FormatUint(uint64(got), Base10)

	err := &ErrTest{
		Kind: kind,
		Want: want_str,
		Got:  got_str,
	}

	return err
}

// Err creates and returns a new ErrTest error with the given expected and
// actual values. Error messages are quoted.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected error.
//   - got: The actual error.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected error message if want is not nil, "no error"
//     otherwise.
//   - <got> is the actual error message if got is not nil, "no error"
//     otherwise.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) Err(kind string, want, got error) error {
	if want == nil && got == nil {
		return nil
	}

	var want_str, got_str string

	if want == nil {
		want_str = "no error"
	} else {
		msg := want.Error()

		want_str = strconv.Quote(msg)
	}

	if got == nil {
		got_str = "no error"
	} else {
		msg := got.Error()

		got_str = strconv.Quote(msg)
	}

	err := &ErrTest{
		Kind: kind,
		Want: want_str,
		Got:  got_str,
	}

	return err
}

// Any creates and returns a new ErrTest error with the given expected and
// actual values.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected value.
//   - got: The actual value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected value if want is not nil, "nothing" otherwise.
//   - <got> is the actual value if got is not nil, "nothing" otherwise.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) Any(kind string, want, got any) error {
	var want_str, got_str string

	if want == nil {
		want_str = ""
	} else {
		want_str = fmt.Sprint(want)
	}

	if got == nil {
		got_str = ""
	} else {
		got_str = fmt.Sprint(got)
	}

	err := &ErrTest{
		Kind: kind,
		Want: want_str,
		Got:  got_str,
	}

	return err
}

// ErrorMessage creates and returns a new ErrTest error with the given expected and
// actual error messages.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected error message.
//   - got: The actual error encountered.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Returns nil if the actual
//     error matches the expected error message, or if both the actual and expected
//     errors are nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected error message if want is not empty, "no error"
//     otherwise.
//   - <got> is the actual error message if got is not nil, "no error" otherwise.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) ErrorMessage(kind, want string, got error) error {
	var want_str, got_str string

	if want == "" {
		want_str = "no error"
	} else {
		want_str = strconv.Quote(want)
	}

	if got == nil {
		got_str = "no error"
	} else {
		msg := got.Error()
		got_str = strconv.Quote(msg)
	}

	err := &ErrTest{
		Kind: kind,
		Want: want_str,
		Got:  got_str,
	}

	return err
}

// Rune creates and returns a new ErrTest error with the given expected and
// actual rune values.
//
// Parameters:
//   - kind: The kind of the value.
//   - want: The expected rune value.
//   - got: The actual rune value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// Format:
//
//	"want <kind> to be <want>, got <got>"
//
// Where:
//   - <kind> is the kind of the value.
//   - <want> is the expected rune value.
//   - <got> is the actual rune value.
//
// If the kind is empty, "want <want>, got <got>" is used.
func (failT) Rune(kind string, want, got rune) error {
	want_str := strconv.QuoteRune(want)
	got_str := strconv.QuoteRune(got)

	err := &ErrTest{
		Kind: kind,
		Want: want_str,
		Got:  got_str,
	}

	return err
}

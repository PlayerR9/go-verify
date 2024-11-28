package test

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// ErrNilReceiver occurs when a method is called on a receiver that is nil. This
	// error can be checked with the == operator.
	//
	// Format:
	// 	"receiver must not be nil"
	ErrNilReceiver error

	// ErrTestNotImpl occurs when a test is not implemented. This can be checked
	// with the == operator.
	//
	// Format:
	// 	"test not implemented"
	ErrTestNotImpl error
)

func init() {
	ErrTestNotImpl = errors.New("test not implemented")
	ErrNilReceiver = errors.New("receiver must not be nil")
}

// ErrTest occurs when a test failed.
type ErrTest struct {
	// Want is the expected value.
	Want string

	// Got is the actual value.
	Got string
}

// Error implements error.
func (e ErrTest) Error() string {
	return "want " + e.Want + ", got " + e.Got
}

// NewErrTest creates and returns a new ErrTest error with the given expected
// and actual values.
//
// Parameters:
//   - want: The expected value.
//   - got: The actual value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest. Never returns nil.
//
// WARNING: Constructors that uses the FAIL namespaces are recommended to be
// used instead of this. However, if the precise use case is not found in the
// FAIL namespaces, this can be used.
//
// Format:
//
//	"want <want>, got <got>"
//
// Where:
//   - <want> is the expected value.
//   - <got> is the actual value.
func NewErrTest(want, got string) error {
	return &ErrTest{
		Want: want,
		Got:  got,
	}
}

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
	var want_str, got_str string

	if want == nil {
		want_str = "no error"
	} else {
		want_str = strconv.Quote(want.Error())
	}

	if got == nil {
		got_str = "no error"
	} else {
		got_str = strconv.Quote(got.Error())
	}

	return &ErrTest{
		Want: want_str,
		Got:  got_str,
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

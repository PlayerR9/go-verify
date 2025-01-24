package test

import (
	"errors"
)

// checkT is for private use only
type checkT struct{}

var (
	// CHECK is the namespace for checking values.
	CHECK checkT = checkT{}
)

// String checks that the given expected and actual string values are equal. If not
// the proper error is returned.
//
// Parameters:
//   - want: The expected string value.
//   - got: The actual string value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest, if the check fails.
func (checkT) String(want, got string) error {
	if want == got {
		return nil
	}

	err := FAIL.String(want, got)
	return err
}

// Int checks that the given expected and actual integer values are equal. If not
// the proper error is returned.
//
// Parameters:
//   - want: The expected integer value.
//   - got: The actual integer value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest, if the check fails.
func (checkT) Int(want, got int) error {
	if want == got {
		return nil
	}

	err := FAIL.Int(want, got)
	return err
}

// Uint checks that the given expected and actual unsigned integer values are
// equal. If not the proper error is returned.
//
// Parameters:
//   - want: The expected unsigned integer value.
//   - got: The actual unsigned integer value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest, if the check fails.
func (checkT) Uint(want, got uint) error {
	if want == got {
		return nil
	}

	err := FAIL.Uint(want, got)
	return err
}

// Err checks that the given expected and actual error values are equal. If not
// the proper error is returned.
//
// Parameters:
//   - want: The expected error value.
//   - got: The actual error value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest, if the check fails.
func (checkT) Err(want, got error) error {
	if want == got {
		return nil
	}

	ok := errors.Is(got, want)
	if ok {
		return nil
	}

	err := FAIL.Err(want, got)
	return err
}

// ErrorMessage checks that the given expected and actual error messages are
// equal. If not the proper error is returned.
//
// Parameters:
//   - want: The expected error message.
//   - got: The actual error message.
//
// Returns:
//   - error: A pointer to the newly created ErrTest, if the check fails.
func (checkT) ErrorMessage(want string, got error) error {
	if want == "" && got == nil {
		return nil
	}

	if want != "" {
		msg := got.Error()
		if msg == want {
			return nil
		}
	}

	err := FAIL.ErrorMessage(want, got)
	return err
}

// Rune checks that the given expected and actual rune values are equal. If not
// the proper error is returned.
//
// Parameters:
//   - want: The expected rune value.
//   - got: The actual rune value.
//
// Returns:
//   - error: A pointer to the newly created ErrTest, if the check fails.
func (checkT) Rune(want, got rune) error {
	if want == got {
		return nil
	}

	err := FAIL.Rune(want, got)
	return err
}

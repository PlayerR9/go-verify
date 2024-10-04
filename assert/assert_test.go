package assert

import (
	"errors"
	"testing"

	test "github.com/PlayerR9/go-verify/test"
)

// MockType is a mock type.
type MockType struct{}

// NewMockType creates a new mock type.
//
// Returns:
//   - *MockType: The new mock type. Always returns nil.
//   - error: Always returns nil.
func NewMockType() (*MockType, error) {
	return nil, nil
}

// IsNil checks whether the mock type is nil.
//
// Returns:
//   - bool: True if the mock type is nil, false otherwise.
func (mt *MockType) IsNil() bool {
	return mt == nil
}

// TestCond tests the Cond function.
func TestCond(t *testing.T) {
	const (
		Expected string = "foo must be bar"
	)

	err := test.CheckPanic(Expected, func() {
		Cond("foo" == "bar", "foo must be bar")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestCondf tests the Condf function.
func TestCondf(t *testing.T) {
	const (
		Expected string = "\"foo\" must be \"bar\""
	)

	err := test.CheckPanic(Expected, func() {
		Condf("foo" == "bar", "%q must be %q", "foo", "bar")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestErr tests the Err function.
func TestErr(t *testing.T) {
	const (
		Expected string = "MyFunc(\"foo\", \"bar\") = something went wrong"
	)

	MyFunc := func(a, b string) error {
		return errors.New("something went wrong")
	}

	err := test.CheckPanic(Expected, func() {
		err := MyFunc("foo", "bar")
		Err(err, "MyFunc(%q, %q)", "foo", "bar")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestTrue tests the True function.
func TestTrue(t *testing.T) {
	const (
		Expected string = "MyFunc(\"foo\", \"bar\") = false"
	)

	MyFunc := func(a, b string) bool {
		return a == b
	}

	err := test.CheckPanic(Expected, func() {
		ok := MyFunc("foo", "bar")
		True(ok, "MyFunc(%q, %q)", "foo", "bar")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestFalse tests the False function.
func TestFalse(t *testing.T) {
	const (
		Expected string = "MyFunc(\"foo\", \"bar\") = true"
	)

	MyFunc := func(a, b string) bool {
		return a != b
	}

	err := test.CheckPanic(Expected, func() {
		ok := MyFunc("foo", "bar")
		False(ok, "MyFunc(%q, %q)", "foo", "bar")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestNotNil tests the NotNil function.
func TestNotNil(t *testing.T) {
	const (
		Expected string = "ms = nil"
	)

	err := test.CheckPanic(Expected, func() {
		var ms *MockType
		NotNil(ms, "ms")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestZero tests the Zero function.
func TestNotZero(t *testing.T) {
	const (
		Expected string = "v = 0"
	)

	err := test.CheckPanic(Expected, func() {
		v := 0
		NotZero(v, "v")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestType tests the Type function.
func TestType(t *testing.T) {
	const (
		Expected string = "v = string, want int"
	)

	err := test.CheckPanic(Expected, func() {
		v := "foo"
		Type[int](v, "v", false)
	})
	if err != nil {
		t.Error(err)
	}
}

// TestDeref tests the Deref function.
func TestDeref(t *testing.T) {
	const (
		Expected string = "v = *int, want string"
	)

	err := test.CheckPanic(Expected, func() {
		var v *int
		_ = Deref[string](v, "v")
	})
	if err != nil {
		t.Error(err)
	}
}

// TestNew tests the New function.
func TestNew(t *testing.T) {
	const (
		Expected string = "*assert.MockType = nil"
	)

	err := test.CheckPanic(Expected, func() {
		New(NewMockType())
	})
	if err != nil {
		t.Error(err)
	}
}

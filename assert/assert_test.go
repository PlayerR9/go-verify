package assert

import (
	"errors"
	"fmt"
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

func TestCond(t *testing.T) {
	type args struct {
		cond     bool
		msg      string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				Cond(args.cond, args.msg)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("cond is true", args{
		cond:     true,
		msg:      "foo",
		expected: "",
	})

	_ = tests.AddTest("cond is false", args{
		cond:     false,
		msg:      "foo",
		expected: "foo",
	})

	_ = tests.Run(t)
}

// TestCondf tests the Condf function.
func TestCondf(t *testing.T) {
	type args struct {
		cond     bool
		format   string
		args     []any
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				Condf(args.cond, args.format, args.args...)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("cond is true", args{
		cond:     true,
		format:   "%q must be %q",
		args:     []any{"foo", "foo"},
		expected: "",
	})

	_ = tests.AddTest("cond is false", args{
		cond:     false,
		format:   "%q must be %q",
		args:     []any{"foo", "bar"},
		expected: fmt.Sprintf("%q must be %q", "foo", "bar"),
	})

	_ = tests.Run(t)
}

// TestErr tests the Err function.
func TestErr(t *testing.T) {
	type args struct {
		inner    error
		format   string
		args     []any
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				Err(args.inner, args.format, args.args...)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("inner is nil", args{
		inner:    nil,
		format:   "%q must be %q",
		args:     []any{"foo", "foo"},
		expected: "",
	})

	_ = tests.AddTest("inner is not nil", args{
		inner:    errors.New("something went wrong"),
		format:   "MyFunc(%q, %q)",
		args:     []any{"foo", "bar"},
		expected: fmt.Sprintf("MyFunc(%q, %q) = %s", "foo", "bar", "something went wrong"),
	})

	_ = tests.Run(t)
}

// TestTrue tests the True function.
func TestTrue(t *testing.T) {
	type args struct {
		ok       bool
		format   string
		args     []any
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				True(args.ok, args.format, args.args...)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("ok is true", args{
		ok:       true,
		format:   "MyFunc(%q, %q)",
		args:     []any{"foo", "foo"},
		expected: "",
	})

	_ = tests.AddTest("ok is false", args{
		ok:       false,
		format:   "MyFunc(%q, %q)",
		args:     []any{"foo", "bar"},
		expected: fmt.Sprintf("MyFunc(%q, %q) = false", "foo", "bar"),
	})

	_ = tests.Run(t)
}

// TestFalse tests the False function.
func TestFalse(t *testing.T) {
	type args struct {
		ok       bool
		format   string
		args     []any
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				False(args.ok, args.format, args.args...)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("ok is false", args{
		ok:       false,
		format:   "MyFunc(%q, %q)",
		args:     []any{"foo", "foo"},
		expected: "",
	})

	_ = tests.AddTest("ok is true", args{
		ok:       true,
		format:   "MyFunc(%q, %q)",
		args:     []any{"foo", "bar"},
		expected: fmt.Sprintf("MyFunc(%q, %q) = true", "foo", "bar"),
	})

	_ = tests.Run(t)
}

// TestNotNil tests the NotNil function.
func TestNotNil(t *testing.T) {
	type args struct {
		v        IsNiler
		name     string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				NotNil(args.v, args.name)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("v is not nil", args{
		v:        &MockType{},
		name:     "v",
		expected: "",
	})

	_ = tests.AddTest("v is nil", args{
		v:        nil,
		name:     "v",
		expected: "v = nil",
	})

	_ = tests.AddTest("v without name", args{
		v:        nil,
		name:     "",
		expected: "variable = nil",
	})

	_ = tests.Run(t)
}

// TestZero tests the Zero function.
func TestNotZero(t *testing.T) {
	type args struct {
		v        int
		name     string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				NotZero(args.v, args.name)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("v is not zero", args{
		v:        1,
		name:     "v",
		expected: "",
	})

	_ = tests.AddTest("v is zero", args{
		v:        0,
		name:     "v",
		expected: "v = 0",
	})

	_ = tests.AddTest("v without name", args{
		v:        0,
		name:     "",
		expected: "variable = 0",
	})

	_ = tests.Run(t)
}

// TestType tests the Type function.
func TestType(t *testing.T) {
	type args struct {
		v         any
		name      string
		allow_nil bool
		expected  string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				Type[int](args.v, args.name, args.allow_nil)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("v is int", args{
		v:         1,
		name:      "v",
		allow_nil: false,
		expected:  "",
	})

	_ = tests.AddTest("v is string", args{
		v:         "foo",
		name:      "v",
		allow_nil: false,
		expected:  "v = string, want int",
	})

	_ = tests.AddTest("v is nil", args{
		v:         nil,
		name:      "v",
		allow_nil: true,
		expected:  "",
	})

	_ = tests.AddTest("v without name", args{
		v:         nil,
		name:      "",
		allow_nil: false,
		expected:  "variable = nil",
	})

	_ = tests.Run(t)
}

// TestDeref tests the Deref function.
func TestDeref(t *testing.T) {
	type args struct {
		v        any
		name     string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				_ = Deref[int](args.v, args.name)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("v is int", args{
		v:        1,
		name:     "v",
		expected: "",
	})

	_ = tests.AddTest("v is string", args{
		v:        "foo",
		name:     "v",
		expected: "v = string, want int",
	})

	_ = tests.AddTest("v without name", args{
		v:        nil,
		name:     "",
		expected: "variable = nil, want int",
	})

	x := 1

	_ = tests.AddTest("v is pointer", args{
		v:        &x,
		name:     "v",
		expected: "",
	})

	_ = tests.Run(t)
}

// TestNew tests the New function.
func TestNew(t *testing.T) {
	type args struct {
		res      *MockType
		inner    error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				New(args.res, args.inner)
			})

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("res is nil", args{
		res:      nil,
		inner:    nil,
		expected: "*assert.MockType = nil",
	})

	_ = tests.AddTest("res is not nil", args{
		res:      &MockType{},
		inner:    nil,
		expected: "",
	})

	_ = tests.AddTest("inner is not nil", args{
		res:      nil,
		inner:    errors.New("foo"),
		expected: "err = foo",
	})

	_ = tests.Run(t)
}

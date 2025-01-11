package assert

import (
	"errors"
	"fmt"
	"testing"

	test "github.com/PlayerR9/go-verify/test"
)

// TestCond tests the Cond function.
func TestCond(t *testing.T) {
	type args struct {
		cond bool
		msg  string
		want string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				Cond(args.cond, args.msg)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("cond is true", args{
		cond: true,
		msg:  "foo",
		want: "",
	})

	_ = tests.Add("cond is false", args{
		cond: false,
		msg:  "foo",
		want: NewErrAssertFail("foo").Error(),
	})

	_ = tests.Run(t)
}

// TestCondf tests the Condf function.
func TestCondf(t *testing.T) {
	type args struct {
		cond   bool
		format string
		args   []any
		want   string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				Condf(args.cond, args.format, args.args...)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("cond is true", args{
		cond:   true,
		format: "%q must be %q",
		args:   []any{"foo", "foo"},
		want:   "",
	})

	_ = tests.Add("cond is false", args{
		cond:   false,
		format: "%q must be %q",
		args:   []any{"foo", "bar"},
		want:   NewErrAssertFail(fmt.Sprintf("%q must be %q", "foo", "bar")).Error(),
	})

	_ = tests.Run(t)
}

// TestErr tests the Err function.
func TestErr(t *testing.T) {
	type args struct {
		inner  error
		format string
		args   []any
		want   string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				Err(args.inner, args.format, args.args...)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("inner is nil", args{
		inner:  nil,
		format: "%q must be %q",
		args:   []any{"foo", "foo"},
		want:   "",
	})

	_ = tests.Add("inner is not nil", args{
		inner:  errors.New("something went wrong"),
		format: "MyFunc(%q, %q)",
		args:   []any{"foo", "bar"},
		want:   NewErrAssertFail(fmt.Sprintf("MyFunc(%q, %q) = %s", "foo", "bar", "something went wrong")).Error(),
	})

	_ = tests.Run(t)
}

// TestTrue tests the True function.
func TestTrue(t *testing.T) {
	type args struct {
		ok     bool
		format string
		args   []any
		want   string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				True(args.ok, args.format, args.args...)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("ok is true", args{
		ok:     true,
		format: "MyFunc(%q, %q)",
		args:   []any{"foo", "foo"},
		want:   "",
	})

	_ = tests.Add("ok is false", args{
		ok:     false,
		format: "MyFunc(%q, %q)",
		args:   []any{"foo", "bar"},
		want:   NewErrAssertFail(fmt.Sprintf("MyFunc(%q, %q) = false", "foo", "bar")).Error(),
	})

	_ = tests.Run(t)
}

// TestFalse tests the False function.
func TestFalse(t *testing.T) {
	type args struct {
		ok     bool
		format string
		args   []any
		want   string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				False(args.ok, args.format, args.args...)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("ok is false", args{
		ok:     false,
		format: "MyFunc(%q, %q)",
		args:   []any{"foo", "foo"},
		want:   "",
	})

	_ = tests.Add("ok is true", args{
		ok:     true,
		format: "MyFunc(%q, %q)",
		args:   []any{"foo", "bar"},
		want:   NewErrAssertFail(fmt.Sprintf("MyFunc(%q, %q) = true", "foo", "bar")).Error(),
	})

	_ = tests.Run(t)
}

// TestZero tests the Zero function.
func TestNotZero(t *testing.T) {
	type args struct {
		v    int
		name string
		want string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				NotZero(args.v, args.name)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("v is not zero", args{
		v:    1,
		name: "v",
		want: "",
	})

	_ = tests.Add("v is zero", args{
		v:    0,
		name: "v",
		want: NewErrAssertFail("v is zero").Error(),
	})

	_ = tests.Add("v without name", args{
		v:    0,
		name: "",
		want: NewErrAssertFail("variable is zero").Error(),
	})

	_ = tests.Run(t)
}

// TestType tests the Type function.
func TestType(t *testing.T) {
	type args struct {
		v         any
		name      string
		allow_nil bool
		want      string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				Type[int](args.v, args.name, args.allow_nil)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("v is int", args{
		v:         1,
		name:      "v",
		allow_nil: false,
		want:      "",
	})

	_ = tests.Add("v is string", args{
		v:         "foo",
		name:      "v",
		allow_nil: false,
		want:      NewErrAssertFail("v = string, want int").Error(),
	})

	_ = tests.Add("v is nil", args{
		v:         nil,
		name:      "v",
		allow_nil: true,
		want:      "",
	})

	_ = tests.Add("v without name", args{
		v:         nil,
		name:      "",
		allow_nil: false,
		want:      NewErrAssertFail("variable = nil").Error(),
	})

	_ = tests.Run(t)
}

// TestDeref tests the Deref function.
func TestDeref(t *testing.T) {
	type args struct {
		v    *int
		name string
		want string
	}

	tests := test.NewTestSet(func(args args) test.TestingFn {
		return func() error {
			caught := test.Try(func() {
				_ = Deref[int](args.v, args.name)
			})

			err := test.FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("v without name", args{
		v:    nil,
		name: "",
		want: NewErrAssertFail("variable = nil, want int").Error(),
	})

	x := 1

	_ = tests.Add("v is pointer", args{
		v:    &x,
		name: "v",
		want: "",
	})

	_ = tests.Run(t)
}

/* // TestNew tests the New function.
func TestNew(t *testing.T) {
	type args struct {
		res      *MockType
		inner    error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func() error {
			err := test.Try(func() {
				New(args.res, args.inner)
			})

			_ = common.FAIL.CheckErr(t, args.expected, err)
		}
	})

	_ = tests.Add("res is nil", args{
		res:      nil,
		inner:    nil,
		expected: NewErrAssertFail("*assert.MockType = nil").Error(),
	})

	_ = tests.Add("res is not nil", args{
		res:      &MockType{},
		inner:    nil,
		expected: "",
	})

	_ = tests.Add("inner is not nil", args{
		res:      nil,
		inner:    errors.New("foo"),
		expected: NewErrAssertFail("err = foo").Error(),
	})

	_ = tests.Run(t)
} */

package test

import (
	"errors"
	"testing"
)

// TestTry tests the Try function and its helper function.
func TestTry(t *testing.T) {
	type args struct {
		fn   func()
		want string
	}

	tests := NewTestSet(func(args args) TestingFn {
		return func() error {
			caught := Try(args.fn)

			err := FAIL.ErrorMessage(caught, args.want)
			return err
		}
	})

	_ = tests.Add("fn does not panic", args{
		fn:   func() {},
		want: "",
	})

	_ = tests.Add("fn panics no error", args{
		fn: func() {
			panic("something went wrong")
		},
		want: "something went wrong",
	})

	_ = tests.Add("fn panics error", args{
		fn: func() {
			panic(errors.New("something went wrong"))
		},
		want: "something went wrong",
	})

	_ = tests.Add("no function", args{
		fn:   nil,
		want: "",
	})

	_ = tests.Run(t)
}

/*
// TestEqualsErr tests the EqualsErr function.
func TestEqualsErr(t *testing.T) {
	type args struct {
		err1 error
		err2 error
		want bool
	}

	tests := NewTestSet(func(args args) TestingFn {
		return func() error {
			err := FAIL.Err(args.err1, args.err2)
			if args.want {
				if err == nil {

				}
			}

			return err
		}
	})

	_ = tests.Add("err self-equal", args{
		err1: errors.New("this is an error"),
		err2: errors.New("this is an error"),
		want: true,
	})

	_ = tests.Add("err equal", args{
		err1: errors.New("this is an error"),
		err2: errors.New("this is an error"),
		want: true,
	})

	_ = tests.Add("err not equal", args{
		err1: errors.New("this is an error"),
		err2: errors.New("this is another error"),
		want: false,
	})

	_ = tests.Run(t)
}
*/

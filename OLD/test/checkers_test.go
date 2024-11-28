package test

import (
	"errors"
	"testing"

	"github.com/PlayerR9/go-verify/OLD/common"
)

// TestTry tests the Try function and its helper function.
func TestTry(t *testing.T) {
	type args struct {
		fn       func()
		expected string
	}

	tests := NewTests(func(args args) TestingFunc {
		return func(t *testing.T) {
			err := Try(args.fn)

			_ = common.FAIL.CheckErr(t, args.expected, err)
		}
	})

	_ = tests.AddTest("fn does not panic", args{
		fn:       func() {},
		expected: "",
	})

	_ = tests.AddTest("fn panics no error", args{
		fn: func() {
			panic("something went wrong")
		},
		expected: "something went wrong",
	})

	_ = tests.AddTest("fn panics error", args{
		fn: func() {
			panic(errors.New("something went wrong"))
		},
		expected: "something went wrong",
	})

	_ = tests.AddTest("no function", args{
		fn:       nil,
		expected: "",
	})

	_ = tests.Run(t)
}

// TestEqualsErr tests the EqualsErr function.
func TestEqualsErr(t *testing.T) {
	type args struct {
		err1     error
		err2     error
		expected bool
	}

	tests := NewTests(func(args args) TestingFunc {
		return func(t *testing.T) {
			ok := EqualsErr(args.err1, args.err2)
			if ok == args.expected {
				return
			}

			common.FAIL.WrongBool(t, args.expected, ok)
		}
	})

	_ = tests.AddTest("err self-equal", args{
		err1:     errors.New("this is an error"),
		err2:     errors.New("this is an error"),
		expected: true,
	})

	_ = tests.AddTest("err equal", args{
		err1:     errors.New("this is an error"),
		err2:     errors.New("this is an error"),
		expected: true,
	})

	_ = tests.AddTest("err not equal", args{
		err1:     errors.New("this is an error"),
		err2:     errors.New("this is another error"),
		expected: false,
	})

	_ = tests.Run(t)
}

package test

import (
	"errors"
	"fmt"
	"testing"
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

			err = CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
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
		expected: "panic: something went wrong",
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
			if ok != args.expected {
				t.Errorf("want %t, got %t", args.expected, ok)
			}
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

// TestCheckErr tests the CheckErr function.
func TestCheckErr(t *testing.T) {
	type args struct {
		expected_str string
		got          error
		expected     string
	}

	tests := NewTests(func(args args) TestingFunc {
		return func(t *testing.T) {
			err := CheckErr(args.expected_str, args.got)
			if args.expected == "" {
				if err != nil {
					t.Errorf("want nil, got %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("want %q, got nil", args.expected)
				} else {
					msg := err.Error()

					if msg != args.expected {
						t.Errorf("want %q, got %q", args.expected, msg)
					}
				}
			}
		}
	})

	_ = tests.AddTest("no error", args{
		expected_str: "",
		got:          nil,
		expected:     "",
	})

	_ = tests.AddTest("no fail", args{
		expected_str: "",
		got:          errors.New("something went wrong"),
		expected:     fmt.Sprintf("want nil, got %q", "something went wrong"),
	})

	_ = tests.AddTest("error", args{
		expected_str: "something went wrong",
		got:          nil,
		expected:     fmt.Sprintf("want %q, got nil", "something went wrong"),
	})

	_ = tests.AddTest("success", args{
		expected_str: "something went wrong",
		got:          errors.New("something went wrong"),
		expected:     "",
	})

	_ = tests.AddTest("fail", args{
		expected_str: "something went wrong",
		got:          errors.New("test failed"),
		expected:     fmt.Sprintf("want %q, got %q", "something went wrong", "test failed"),
	})

	_ = tests.Run(t)
}

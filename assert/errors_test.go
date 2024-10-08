package assert

import (
	"errors"
	"testing"

	test "github.com/PlayerR9/go-verify/test"
)

// TestNewErrAssertFailed tests the NewErrAssertFailed function.
func TestNewErrAssertFailed(t *testing.T) {
	type args struct {
		message  string
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrAssertFailed(args.message)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with message", args{
		message:  "foo",
		expected: "foo",
	})

	_ = tests.AddTest("without message", args{
		message:  "",
		expected: "something went wrong",
	})

	_ = tests.Run(t)
}

// TestNewErrValidationFailed tests the NewErrValidationFailed function.
func TestNewErrValidationFailed(t *testing.T) {
	type args struct {
		name     string
		reason   error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrValidationFailed(args.name, args.reason)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with name", args{
		name:     "foo",
		reason:   nil,
		expected: "foo = nil",
	})

	_ = tests.AddTest("with reason", args{
		name:     "",
		reason:   errors.New("test reason"),
		expected: "struct = test reason",
	})

	_ = tests.Run(t)
}

// TestNewErrFixFailed tests the NewErrFixFailed function.
func TestNewErrFixFailed(t *testing.T) {
	type args struct {
		name     string
		reason   error
		expected string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := NewErrFixFailed(args.name, args.reason)

			err = test.CheckErr(args.expected, err)
			if err != nil {
				t.Error(err)
			}
		}
	})

	_ = tests.AddTest("with name", args{
		name:     "foo",
		reason:   nil,
		expected: "foo = nil",
	})

	_ = tests.AddTest("with reason", args{
		name:     "",
		reason:   errors.New("test reason"),
		expected: "struct = test reason",
	})

	_ = tests.Run(t)
}

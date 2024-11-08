package assert

import (
	"errors"
	"testing"

	"github.com/PlayerR9/go-verify/common"
	"github.com/PlayerR9/go-verify/test"
)

// MockStruct is a mock struct.
type MockStruct struct {
	// name is the name of the mock struct.
	name string
}

// Validate implements Validater.
func (ms MockStruct) Validate() error {
	if ms.name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

// Fix implements Fixer.
func (ms *MockStruct) Fix() error {
	if ms == nil {
		return nil
	}

	if ms.name == "" {
		return errors.New("name cannot be empty")
	}

	ms.name = "foo"

	return nil
}

// TestValidate tests the Validate function.
func TestValidate(t *testing.T) {
	type args struct {
		v         Validater
		name      string
		allow_nil bool
		expected  string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				Validate(args.v, args.name, args.allow_nil)
			})

			_ = common.FAIL.CheckErr(t, args.expected, err)
		}
	})

	_ = tests.AddTest("valid", args{
		v: &MockStruct{
			name: "foo",
		},
		name:      "ms",
		allow_nil: false,
		expected:  "",
	})

	_ = tests.AddTest("nil without allow_nil", args{
		v:         nil,
		name:      "",
		allow_nil: false,
		expected:  "(Validate Failed) struct = nil",
	})

	_ = tests.AddTest("nil with allow_nil", args{
		v:         nil,
		name:      "",
		allow_nil: true,
		expected:  "",
	})

	_ = tests.AddTest("invalid", args{
		v: &MockStruct{
			name: "",
		},
		name:      "ms",
		allow_nil: false,
		expected:  "(Validate Failed) ms = name cannot be empty",
	})

	_ = tests.Run(t)
}

// TestFix tests the Fix function.
func TestFix(t *testing.T) {
	type args struct {
		v         Fixer
		name      string
		allow_nil bool
		expected  string
	}

	tests := test.NewTests(func(args args) test.TestingFunc {
		return func(t *testing.T) {
			err := test.Try(func() {
				Fix(args.v, args.name, args.allow_nil)
			})

			_ = common.FAIL.CheckErr(t, args.expected, err)
		}
	})

	_ = tests.AddTest("valid", args{
		v: &MockStruct{
			name: "foo",
		},
		name:      "ms",
		allow_nil: false,
		expected:  "",
	})

	_ = tests.AddTest("nil without allow_nil", args{
		v:         nil,
		name:      "",
		allow_nil: false,
		expected:  "(Fix Failed) struct = nil",
	})

	_ = tests.AddTest("nil with allow_nil", args{
		v:         nil,
		name:      "",
		allow_nil: true,
		expected:  "",
	})

	_ = tests.AddTest("invalid", args{
		v: &MockStruct{
			name: "",
		},
		name:      "ms",
		allow_nil: false,
		expected:  "(Fix Failed) ms = name cannot be empty",
	})

	_ = tests.Run(t)
}

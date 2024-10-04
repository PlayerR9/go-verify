package assert

import (
	"errors"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// MockType is a mock type.
type MockStruct struct {
	// Name is a mock field.
	Name string
}

// Validate implements the Validater interface.
//
// An instance of MockStruct is valid iff:
//   - The field "Name" is not empty.
func (ms MockStruct) Validate() error {
	if ms.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

// Fix implements the Fixer interface.
//
// An instance of MockStruct is valid iff:
//   - The field "Name" is not empty.
func (ms *MockStruct) Fix() error {
	if ms == nil {
		return nil
	}

	if ms.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

// TestValidate tests the Validate function.
func TestValidate(t *testing.T) {
	const (
		Expected string = "ms = name cannot be empty"
	)

	ms := &MockStruct{
		Name: "",
	}

	err := test.CheckPanic(Expected, func() {
		Validate(ms, "ms", false)
	})
	if err != nil {
		t.Error(err)
	}
}

// TestFix tests the Fix function.
func TestFix(t *testing.T) {
	const (
		Expected string = "ms = name cannot be empty"
	)

	ms := &MockStruct{}

	err := test.CheckPanic(Expected, func() {
		Fix(ms, "ms", false)
	})
	if err != nil {
		t.Error(err)
	}
}

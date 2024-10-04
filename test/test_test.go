package test

import (
	"errors"
	"testing"
)

// TestCheckPanic tests the CheckPanic function.
func TestCheckPanic(t *testing.T) {
	const (
		Expected string = "something went wrong"
	)

	err := CheckPanic(Expected, func() {
		panic(errors.New(Expected))
	})
	if err != nil {
		t.Error(err)
	}
}

// TestCheckErr tests the CheckErr function.
func TestCheckErr(t *testing.T) {
	const (
		Expected string = "something went wrong"
	)

	err := CheckErr(Expected, errors.New(Expected))
	if err != nil {
		t.Error(err)
	}
}

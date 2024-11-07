package test

import "testing"

// TestingFunc is the type of the testing function.
//
// Parameters:
//   - t: The testing instance.
type TestingFunc func(t *testing.T)

var (
	// UnimplementedTest is the error returned when a test is not implemented.
	UnimplementedTest TestingFunc
)

func init() {
	UnimplementedTest = func(t *testing.T) {
		t.Error("test not implemented")
	}
}

// MakeTestFn is a function that creates a test function.
//
// Parameters:
//   - args: The arguments passed to the test function.
//
// Returns:
//   - TestingFunc: The test function. Never returns nil.
type MakeTestFn[T any] func(args T) TestingFunc

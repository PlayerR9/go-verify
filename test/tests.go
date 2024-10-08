package test

import "testing"

// TestingFunc is the type of the testing function.
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

// Instance is a test instance.
type Instance[T any] struct {
	// name is the name of the test.
	name string

	// args are the arguments passed to the test function.
	args T

	// fn is the test function.
	fn TestingFunc
}

// Tests is a collection of test instances.
type Tests[T any] struct {
	// tests is the collection of test instances.
	tests []*Instance[T]

	// make_test is the function that creates the test function.
	make_test func(args T) TestingFunc
}

// NewTests creates a new test collection.
//
// Parameters:
//   - make_test: The function that creates the test function. If nil, the
//     UnimplementedTest function is used.
//
// Returns:
//   - Tests[T]: The new test collection.
func NewTests[T any](make_test func(args T) TestingFunc) Tests[T] {
	if make_test == nil {
		make_test = func(args T) TestingFunc {
			return UnimplementedTest
		}
	}

	return Tests[T]{
		tests:     make([]*Instance[T], 0),
		make_test: make_test,
	}
}

// AddTest adds a new test to the collection.
//
// Parameters:
//   - name: The name of the test.
//   - args: The arguments passed to the test function.
//
// Returns:
//   - bool: True if the receiver is not nil, false otherwise.
func (t *Tests[T]) AddTest(name string, args T) bool {
	if t == nil {
		return false
	}

	test := &Instance[T]{
		name: name,
		args: args,
		fn:   t.make_test(args),
	}

	t.tests = append(t.tests, test)

	return true
}

// Run runs all the tests in the collection. They are run in the same order as
// they were added.
//
// Parameters:
//   - t: The testing object. If nil, no tests will be run.
//
// Returns:
//   - int: The number of tests that passed.
func (tests Tests[T]) Run(t *testing.T) int {
	if t == nil {
		return 0
	}

	var count int

	for _, test := range tests.tests {
		ok := t.Run(test.name, test.fn)
		if ok {
			count++
		}
	}

	return count
}

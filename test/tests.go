package test

import (
	"testing"

	"github.com/PlayerR9/go-verify/common"
)

// Tests is a collection of test instances.
type Tests[T any] struct {
	// tests is the collection of test instances.
	tests []Instance[T]

	// make_test is the function that creates the test function.
	make_test MakeTestFn[T]
}

// NewTests creates a new test collection.
//
// Parameters:
//   - make_test: The function that creates the test function. If nil, the
//     UnimplementedTest function is used.
//
// Returns:
//   - Tests[T]: The new test collection. Never returns nil.
func NewTests[T any](make_test MakeTestFn[T]) Tests[T] {
	if make_test == nil {
		make_test = func(args T) TestingFunc {
			return UnimplementedTest
		}
	}

	return Tests[T]{
		tests:     make([]Instance[T], 0),
		make_test: make_test,
	}
}

// GetTestsCount returns the number of tests in the collection.
//
// Returns:
//   - uint: The number of tests in the collection.
func (tests Tests[T]) GetTestsCount() uint {
	return uint(len(tests.tests))
}

// AddTest adds a new test to the collection.
//
// Parameters:
//   - name: The name of the test.
//   - args: The arguments passed to the test function.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (tests *Tests[T]) AddTest(name string, args T) error {
	if tests == nil {
		return common.ErrNilReceiver
	}

	test := Instance[T]{
		name: name,
		args: args,
		fn:   tests.make_test(args),
	}

	tests.tests = append(tests.tests, test)

	return nil
}

// Run runs all the tests in the collection. These tests are run in the same
// order as they were added and, because it uses the `Run()` method of the
// testing object, they are run in parallel.
//
// Does nothing if no tests were added.
//
// Parameters:
//   - t: The testing object.
//
// Returns:
//   - uint: The number of tests that passed.
//
// Panics:
//   - NoTestInstance: If the testing object is nil.
func (tests Tests[T]) Run(t *testing.T) uint {
	if len(tests.tests) == 0 {
		return 0
	} else if t == nil {
		panic(common.NoTestInstance)
	}

	var count uint

	for _, test := range tests.tests {
		ok := t.Run(test.name, test.fn)
		if !ok {
			continue
		}

		count++
	}

	return count
}

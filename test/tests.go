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
//   - Tests[T]: The new test collection.
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

// AddTest adds a new test to the collection.
//
// Parameters:
//   - name: The name of the test.
//   - args: The arguments passed to the test function.
//
// Returns:
//   - error: An error of type common.ErrNilReceiver if the receiver is nil.
func (t *Tests[T]) AddTest(name string, args T) error {
	if t == nil {
		return common.ErrNilReceiver
	}

	test := Instance[T]{
		name: name,
		args: args,
		fn:   t.make_test(args),
	}

	t.tests = append(t.tests, test)

	return nil
}

// Run runs all the tests in the collection. They are run in the same order as
// they were added. However, they are run in parallel.
//
// Parameters:
//   - t: The testing object. If nil, no tests will be run.
//
// Returns:
//   - int: The number of tests that passed.
func (tests Tests[T]) Run(t *testing.T) int {
	if len(tests.tests) == 0 {
		return 0
	} else if t == nil {
		panic(NoTestInstance)
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

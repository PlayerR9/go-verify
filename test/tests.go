package test

import (
	"testing"
)

// TestSet is a collection of tests.
type TestSet[T any] struct {
	// makeFn is a function that is used to create TestingFn instances.
	makeFn MakeFn[T]

	// instances is the collection of tests.
	instances []Instance
}

// NewTestSet creates and returns a new Tests instance with a specified
// function to create TestingFn instances. If no function is provided,
// a default testing function is used.
//
// Parameters:
//   - makeFn: A function that returns a TestingFn. If nil, a default
//     testing function is used.
//
// Returns:
//   - Tests: A new Tests instance with the provided or default makeFn.
func NewTestSet[T any](makeFn MakeFn[T]) TestSet[T] {
	if makeFn == nil {
		makeFn = func(_ T) TestingFn {
			return DefaultTestingFn
		}
	}

	return TestSet[T]{
		makeFn: makeFn,
	}
}

// Add adds a new test to the collection of tests.
//
// Parameters:
//   - name: The name of the test.
//   - args: The arguments to pass to the testing function.
//
// Returns:
//   - error: An error if the test could not be added.
//
// Errors:
//   - ErrNilReceiver: If the receiver is nil.
func (tt *TestSet[T]) Add(name string, args T) error {
	if tt == nil {
		return ErrNilReceiver
	}

	instance := Instance{
		name: name,
		fn:   tt.makeFn(args),
	}

	tt.instances = append(tt.instances, instance)

	return nil
}

// Run runs all tests in the collection. Does nothing if there are no tests.
//
// Parameters:
//   - t: The testing.T instance to use for reporting.
//
// Returns:
//   - uint: The number of tests that are failed.
//
// Panics:
//   - "parameter (t) must not be nil": If t is nil.
func (tt TestSet[T]) Run(t *testing.T) uint {
	if len(tt.instances) == 0 {
		return 0
	} else if t == nil {
		panic("parameter (t) must not be nil")
	}

	var count uint

	for _, instance := range tt.instances {
		fn := func(t *testing.T) {
			err := instance.fn()
			if err == nil {
				return
			}

			t.Error(err)
		}

		ok := t.Run(instance.name, fn)
		if !ok {
			count++
		}
	}

	return count
}

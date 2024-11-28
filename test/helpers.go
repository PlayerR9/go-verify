package test

// TestingFn is a function that is used to run a test.
//
// Parameters:
//   - error: An error if the test failed.
type TestingFn func() error

var (
	// DefaultTestingFn is the default function for when no testing function is provided.
	//
	// It just returns ErrTestNotImpl.
	DefaultTestingFn TestingFn
)

func init() {
	DefaultTestingFn = func() error {
		return ErrTestNotImpl
	}
}

// MakeFn is a function that is used to create TestingFn instances.
//
// Parameters:
//   - args: The arguments to pass to the testing function.
//
// Returns:
//   - TestingFn: The function that is used to run the test. Never returns nil.
type MakeFn[T any] func(args T) TestingFn

// Instance is an instance of a test.
type Instance struct {
	// name is the name of the test.
	name string

	// fn is the function that is used to run the test.
	fn TestingFn
}

package test

// Instance is a test instance.
type Instance[T any] struct {
	// name is the name of the test.
	name string

	// args are the arguments passed to the test function.
	args T

	// fn is the test function.
	fn TestingFunc
}

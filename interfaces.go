package errs

// ErrorCloner is the interface that wraps the CloneError method.
type ErrorCloner interface {
	// CloneError clones the error in a shallow way.
	//
	// Returns:
	//   - error: A pointer to a new error. Never returns nil.
	CloneError() error
}

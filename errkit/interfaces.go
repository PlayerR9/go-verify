package errkit

import "github.com/PlayerR9/go-verify/errkit/internal"

// ErrorCloner is the interface that wraps the CloneError method.
type ErrorCloner interface {
	// CloneError clones the error in a shallow way.
	//
	// Returns:
	//   - ErrorCloner: A pointer to a new error. Never returns nil.
	CloneError() ErrorCloner

	// GetInfo returns the info of the error.
	//
	// Returns:
	//   - *Info: The info of the error.
	GetInfo() *internal.Info

	// SetInfo sets the info of the error.
	//
	// Parameters:
	//   - info: The info to set.
	SetInfo(info *internal.Info)

	error
}

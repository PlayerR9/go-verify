package common

import "errors"

var (
	// ErrInvalidBase occurs when the base is zero.
	//
	// Format:
	//   "base must not be zero"
	ErrInvalidBase error

	// ErrNilReceiver occurs when the receiver is nil.
	//
	// Format:
	//   "receiver must not be nil"
	ErrNilReceiver error
)

func init() {
	ErrInvalidBase = errors.New("base must not be zero")

	ErrNilReceiver = errors.New("receiver must not be nil")
}

package errkit

import (
	"fmt"
	"io"

	"github.com/PlayerR9/go-verify/errkit/internal"
)

// ErrorCoder is the interface representing the error code.
type ErrorCoder interface {
	~int

	fmt.Stringer
}

// CodedErr is a custom error type for complex error information.
type CodedErr[C ErrorCoder] struct {
	// Level is the level of the error.
	Level ErrorLevel

	// Code is the code of the error.
	Code C

	// Msg is the reason of the error.
	Msg string

	// Info is the info of the error.
	*internal.Info
}

// Error implements the error interface.
//
// Format:
//
//	"[<level>] <code>: <msg>"
//
// where:
//   - <level>: The level of the error.
//   - <code>: The error code. If nil, "unknown" is used.
//   - <msg>: The error message. If empty, "something went wrong" is used.
func (e CodedErr[C]) Error() string {
	var msg string

	if e.Msg == "" {
		msg = "something went wrong"
	} else {
		msg = e.Msg
	}

	return "[" + e.Level.String() + "] " + e.Code.String() + ": " + msg
}

// WriteInfo implements the InfoWriter interface.
func (e CodedErr[C]) WriteInfo(w io.Writer) error {
	return write_info(w, e.Info)
}

// New creates a new CodedErr with the given error code and message.
//
// Parameters:
//   - code: The error code.
//   - msg: The error message.
//
// Returns:
//   - *CodedErr: The new error. Never returns nil.
func New[C ErrorCoder](code C, msg string) *CodedErr[C] {
	return &CodedErr[C]{
		Level: ERROR,
		Code:  code,
		Msg:   msg,
		Info:  internal.NewInfo(),
	}
}

// NewWithSeverity creates a new CodedErr with the given severity, error code and message.
//
// Parameters:
//   - severity: The severity of the error.
//   - code: The error code.
//   - msg: The error message.
//
// Returns:
//   - *CodedErr: The new error. Never returns nil.
func NewWithSeverity[C ErrorCoder](severity ErrorLevel, code C, msg string) *CodedErr[C] {
	return &CodedErr[C]{
		Level: severity,
		Code:  code,
		Msg:   msg,
		Info:  internal.NewInfo(),
	}
}

// NewFromError creates a new error from an error.
//
// Parameters:
//   - code: The error code.
//   - err: The error to wrap.
//
// Returns:
//   - *Err: A pointer to the new error. Never returns nil.
func NewFromError[C ErrorCoder](code C, err error) *CodedErr[C] {
	var outer *CodedErr[C]

	if err == nil {
		outer = &CodedErr[C]{
			Code: code,
			Msg:  "something went wrong",
			Info: internal.NewInfo(),
		}
	} else {
		switch inner := err.(type) {
		case interface {
			ClearInfo()
			GetMessage() string
			Copy() *internal.Info
		}:
			outer = &CodedErr[C]{
				Code: code,
				Msg:  inner.GetMessage(),
				Info: inner.Copy(),
			}

			inner.ClearInfo() // Clear any info since it is now in the outer error.
		default:
			outer = &CodedErr[C]{
				Code: code,
				Msg:  inner.Error(),
				Info: internal.NewInfo(),
			}
		}
	}

	outer.Level = ERROR

	return outer
}

// ClearInfo clears the info of the error. Does nothing
// if the receiver is nil.
func (e *CodedErr[C]) ClearInfo() {
	if e == nil {
		return
	}

	e.Info = nil
}

// CloneError clones the error in a shallow way.
//
// Returns:
//   - error: A pointer to a new error. Never returns nil.
func (ce CodedErr[C]) CloneError() error {
	return &CodedErr[C]{
		Level: ce.Level,
		Code:  ce.Code,
		Msg:   ce.Msg,
		Info:  ce.Info,
	}
}

// GetMessage gets the message of the error.
//
// Returns:
//   - string: The message of the error.
func (e CodedErr[C]) GetMessage() string {
	return e.Msg
}

// IsNil checks if the error is nil.
//
// Returns:
//   - bool: True if the error is nil, false otherwise.
func (e *CodedErr[C]) IsNil() bool {
	return e == nil
}

// ChangeSeverity changes the severity of the error. Does nothing
// if the receiver is nil.
//
// Parameters:
//   - new_level: The new severity of the error.
func (e *CodedErr[C]) ChangeSeverity(new_level ErrorLevel) {
	if e == nil {
		return
	}

	e.Level = new_level
}

// ModifyLevel modifies the level of the error. Does nothing
// if the receiver is nil.
//
// Parameters:
//   - new_level: The new level of the error.
func (e *CodedErr[C]) ModifyLevel(new_level ErrorLevel) {
	if e == nil {
		return
	}

	e.Level = new_level
}

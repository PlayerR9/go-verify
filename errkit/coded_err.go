package errkit

import (
	"fmt"
	"io"
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

	// Context is the context of the error.
	Context map[string]any

	// Suggestions are suggestions for the user.
	Suggestions []string
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
	if e.Context != nil {
		err := WriteString(w, "Context:\n")
		if err != nil {
			return err
		}

		for key, value := range e.Context {
			_, err = fmt.Fprintf(w, "- %s: %v\n", key, value)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
		Level:   ERROR,
		Code:    code,
		Msg:     msg,
		Context: make(map[string]any),
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
		Level:   severity,
		Code:    code,
		Msg:     msg,
		Context: make(map[string]any),
	}
}

// IsNil checks if the error is nil.
//
// Returns:
//   - bool: True if the error is nil, false otherwise.
func (e *CodedErr[C]) IsNil() bool {
	return e == nil
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

// Add adds a key-value pair to the set. Overwrites any
// existing value if the key already exists.
//
// Parameters:
//   - key: The key of the pair.
//   - value: The value of the pair.
//
// Returns:
//   - bool: True if the key was added, false otherwise.
func (e *CodedErr[C]) AddContext(key string, value any) bool {
	if e == nil {
		return false
	}

	e.Context[key] = value

	return true
}

// Get gets a key from the set.
//
// Returns:
//   - any: The value of the key, or the zero value if the key does not exist.
//   - bool: True if the key exists, false otherwise.
func (e CodedErr[C]) Get(key string) (any, bool) {
	val, ok := e.Context[key]
	return val, ok
}

// Has checks if the set has a key.
//
// Returns:
//   - bool: True if the key exists, false otherwise.
func (e CodedErr[C]) Has(key string) bool {
	_, ok := e.Context[key]
	return ok
}

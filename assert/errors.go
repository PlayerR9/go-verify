package assert

// ErrAssertFailed is an error that is returned when an assertion fails.
type ErrAssertFailed struct {
	// Message describes what went wrong.
	Message string
}

// Error implements the error interface.
//
// Format:
//
//	"<msg>"
//
// where <msg> is the message. If empty, "something went wrong" is used.
func (e ErrAssertFailed) Error() string {
	var msg string

	if e.Message == "" {
		msg = "something went wrong"
	} else {
		msg = e.Message
	}

	return msg
}

// NewErrAssertFailed creates a new ErrAssertFailed.
//
// Parameters:
//   - reason: The message of the error.
//
// Returns:
//   - *ErrAssertFailed: The new error. Never returns nil.
func NewErrAssertFailed(message string) *ErrAssertFailed {
	return &ErrAssertFailed{
		Message: message,
	}
}

// IsNil checks if the error is nil.
//
// Returns:
//   - bool: True if the error is nil, false otherwise.
func (e *ErrAssertFailed) IsNil() bool {
	return e == nil
}

// ErrValidationFailed is an error that is returned when a validation fails.
type ErrValidationFailed struct {
	// Name is the name of the variable. If empty, "variable" is used.
	Name string

	// Reason describes what went wrong.
	Reason error
}

// Error implements the error interface.
//
// Format:
//
//	"<name> = <reason>"
//
// where:
//   - <name> is the name of the variable. If empty, "struct" is used.
//   - <reason> is the message of the error.
func (e ErrValidationFailed) Error() string {
	var name string

	if e.Name == "" {
		name = "struct"
	} else {
		name = e.Name
	}

	var msg string

	if e.Reason == nil {
		msg = name + " = nil"
	} else {
		msg = name + " = " + e.Reason.Error()
	}

	return msg
}

// NewErrValidationFailed creates a new ErrValidationFailed.
//
// Parameters:
//   - name: The name of the variable.
//   - reason: The message of the error.
//
// Returns:
//   - *ErrValidationFailed: The new error. Never returns nil.
func NewErrValidationFailed(name string, reason error) *ErrValidationFailed {
	return &ErrValidationFailed{
		Name:   name,
		Reason: reason,
	}
}

// IsNil checks if the error is nil.
//
// Returns:
//   - bool: True if the error is nil, false otherwise.
func (e *ErrValidationFailed) IsNil() bool {
	return e == nil
}

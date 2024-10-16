package assert

// ErrAssertFailed is an error that is returned when an assertion fails.
type ErrAssertFailed struct {
	// Message describes what went wrong.
	Message string
}

// Error implements the error interface.
func (e ErrAssertFailed) Error() string {
	var msg string

	if e.Message == "" {
		msg = "something went wrong"
	} else {
		msg = e.Message
	}

	return "(Assertion Failed): " + msg
}

// NewErrAssertFailed creates a new ErrAssertFailed.
//
// Parameters:
//   - reason: The message of the error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"<msg>"
//
// where <msg> is the message. If empty, "something went wrong" is used.
func NewErrAssertFailed(message string) error {
	return &ErrAssertFailed{
		Message: message,
	}
}

// ErrValidationFailed is an error that is returned when a validation fails.
type ErrValidationFailed struct {
	// Name is the name of the variable. If empty, "variable" is used.
	Name string

	// Reason describes what went wrong.
	Reason error
}

// Error implements the error interface.
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
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"<name> = <reason>"
//
// where:
//   - <name> is the name of the variable. If empty, "struct" is used.
//   - <reason> is the message of the error.
func NewErrValidationFailed(name string, reason error) error {
	return &ErrValidationFailed{
		Name:   name,
		Reason: reason,
	}
}

// ErrFixFailed is an error that is returned when a fix fails.
type ErrFixFailed struct {
	// Name is the name of the variable. If empty, "variable" is used.
	Name string

	// Reason describes what went wrong.
	Reason error
}

// Error implements the error interface.
func (e ErrFixFailed) Error() string {
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

// NewErrFixFailed creates a new ErrFixFailed.
//
// Parameters:
//   - name: The name of the variable.
//   - reason: The message of the error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	"<msg>"
//
// where <msg> is the message. If empty, "something went wrong" is used.
func NewErrFixFailed(name string, reason error) error {
	return &ErrFixFailed{
		Name:   name,
		Reason: reason,
	}
}

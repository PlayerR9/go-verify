package assert

// ErrAssertFailed is an error that is returned when an assertion fails.
type ErrAssertFailed struct {
	// Msg describes what went wrong.
	Msg string
}

// Error implements the error interface.
func (e ErrAssertFailed) Error() string {
	var msg string

	if e.Msg == "" {
		msg = "something went wrong"
	} else {
		msg = e.Msg
	}

	return "(Assertion Failed) " + msg
}

// NewErrAssertFailed creates a new ErrAssertFailed.
//
// Parameters:
//   - msg: The message of the error.
//
// Returns:
//   - error: The new error. Never returns nil.
//
// Format:
//
//	(Assertion Failed) "<msg>"
//
// where <msg> is the message. If empty, "something went wrong" is used.
func NewErrAssertFailed(msg string) error {
	return &ErrAssertFailed{
		Msg: msg,
	}
}

// ErrValidateFailed is an error that is returned when a validation fails.
type ErrValidateFailed struct {
	// Name is the name of the variable. If empty, "variable" is used.
	Name string

	// Reason describes what went wrong.
	Reason error
}

// Error implements the error interface.
func (e ErrValidateFailed) Error() string {
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

	return "(Validate Failed) " + msg
}

// NewErrValidateFailed creates a new ErrValidateFailed.
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
//	(Validate Failed) "<name> = <reason>"
//
// where:
//   - <name> is the name of the variable. If empty, "struct" is used.
//   - <reason> is the message of the error.
func NewErrValidateFailed(name string, reason error) error {
	return &ErrValidateFailed{
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

	return "(Fix Failed) " + msg
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
//	(Fix Failed) "<msg>"
//
// where <msg> is the message. If empty, "something went wrong" is used.
func NewErrFixFailed(name string, reason error) error {
	return &ErrFixFailed{
		Name:   name,
		Reason: reason,
	}
}

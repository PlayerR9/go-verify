package assert

// ErrAssertFail occurs when an assertion is not met.
type ErrAssertFail struct {
	// Msg is the error message.
	Msg string
}

// Error implements error.
func (e ErrAssertFail) Error() string {
	var msg string

	if e.Msg != "" {
		msg = e.Msg
	} else {
		msg = "an assertion was not met"
	}

	return "[ASSERT FAIL]: " + msg
}

// NewErrAssertFail creates and returns a new ErrAssertFail error with the
// specified error message.
//
// Parameters:
//   - msg: The error message.
//
// Returns:
//   - error: A pointer to the newly created ErrAssertFail. Never returns nil.
//
// Format:
//
//	"[ASSERT FAIL]: <msg>"
//
// Where:
//   - <msg> is the error message. If empty, defaults to "an assertion was not
//     met".
func NewErrAssertFail(msg string) error {
	err := &ErrAssertFail{
		Msg: msg,
	}

	return err
}

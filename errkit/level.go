package errkit

// ErrorLevel is the level of an error. The lower the level, the more severe the error.
type ErrorLevel int

const (
	// _Unknown is used for when the error level is unknown.
	_Unknown ErrorLevel = iota - 1

	// FATAL is the error level used for panic-level type of errors.
	FATAL

	// ERROR is the standard error level. Used for most errors.
	ERROR

	// WARNING is not a critical error. Used for warnings and/or non-fatal errors.
	WARNING

	// DEBUG is used mainly during development. Not used in production.
	DEBUG

	// INFO is used for informational messages. Mostly used for message-passing
	// and channel messages.
	INFO
)

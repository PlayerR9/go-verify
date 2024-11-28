package assert

import (
	"io"
	"os"
)

// TODO panics with a TODO message. The given message is appended to the
// string "TODO: ". If the message is empty, the message "TODO: Handle this
// case" is used instead.
//
// Parameters:
//   - msg: The message to append to the string "TODO: ".
//
// This function is meant to be used only when the code is being built or
// refactored.
func TODO(msg string) {
	if msg == "" {
		panic("TODO: Handle this case")
	} else {
		panic("TODO: " + msg)
	}
}

// WARN prints a warning message to the console.
// The message is prefixed with "[WARNING]:" to indicate its nature.
//
// Parameters:
//   - msg: The warning message to be displayed.
//
// Panics if there is an error writing to the standard output.
func WARN(msg string) {
	data := []byte("[WARNING]: " + msg + "\n")

	n, err := os.Stdout.Write(data)
	if err != nil {
		panic(err)
	} else if n != len(data) {
		panic(io.ErrShortWrite)
	}
}

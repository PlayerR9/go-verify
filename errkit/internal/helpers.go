package internal

import (
	"fmt"
	"io"
)

// WriteString writes the specified string in the writer, if it is not empty.
//
// Parameters:
//   - w: The writer to write to.
//   - str: The string to write.
//
// Returns:
//   - error: If the string could not be written in full.
func WriteString(w io.Writer, str string) error {
	if str == "" {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(str)

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}

// WriteStringf writes the formatted string in the writer, if it is not empty.
//
// Parameters:
//   - w: The writer to write to.
//   - format: The format of the string.
//   - args: The arguments of the format.
//
// Returns:
//   - error: If the string could not be written in full.
func WriteStringf(w io.Writer, format string, args ...any) error {
	str := fmt.Sprintf(format, args...)

	if str == "" {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(str)

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	return nil
}

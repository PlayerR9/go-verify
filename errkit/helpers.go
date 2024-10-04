package errkit

import "io"

// InfoWriter is the interface that wraps the WriteInfo method.
type InfoWriter interface {
	// WriteInfo writes the error in full.
	//
	// Parameters:
	//   - w: The writer to write to.
	//
	// Returns:
	//   - error: If the error could not be displayed in full.
	WriteInfo(w io.Writer) error
}

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

// DisplayError displays the error in full. It prints the error message and, if any,
// optional information.
//
// Parameters:
//   - w: The writer to write to.
//   - to_display: The error to display.
//
// Returns:
//   - error: If the error could not be displayed in full.
func DisplayError(w io.Writer, to_display error) error {
	if to_display == nil {
		return nil
	} else if w == nil {
		return io.ErrShortWrite
	}

	data := []byte(to_display.Error())

	n, err := w.Write(data)
	if err != nil {
		return err
	} else if n != len(data) {
		return io.ErrShortWrite
	}

	info, ok := to_display.(InfoWriter)
	if !ok {
		return nil
	}

	err = info.WriteInfo(w)
	return err
}

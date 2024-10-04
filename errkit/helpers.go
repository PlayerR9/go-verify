package errkit

import (
	"io"
	"slices"
	"strings"

	"github.com/PlayerR9/go-verify/errkit/internal"
)

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

// write_info displays the info to the writer.
//
// Parameters:
//   - w: The writer to write to.
//   - info: The info to display.
//
// Returns:
//   - error: The error that occurred while displaying the info.
func write_info(w io.Writer, info *internal.Info) error {
	if info == nil {
		return nil
	}

	if !info.Timestamp.IsZero() {
		err := internal.WriteStringf(w, "Occurred at: %v\n", info.Timestamp)
		if err != nil {
			return err
		}
	}

	if len(info.Suggestions) > 0 {
		err := internal.WriteString(w, "Suggestions:\n")
		if err != nil {
			return err
		}

		for _, suggestion := range info.Suggestions {
			err := internal.WriteStringf(w, "- %s\n", suggestion)
			if err != nil {
				return err
			}
		}
	}

	if len(info.Context) > 0 {
		err := internal.WriteString(w, "\nContext:\n")
		if err != nil {
			return err
		}

		for k, v := range info.Context {
			err := internal.WriteStringf(w, "- %s: %v\n", k, v)
			if err != nil {
				return err
			}
		}
	}

	if info.Frames != nil {
		err := internal.WriteString(w, "\nStack trace:\n")
		if err != nil {
			return err
		}

		elem := make([]string, len(info.Frames))
		copy(elem, info.Frames)

		slices.Reverse(elem)

		err = internal.WriteStringf(w, "- %s\n", strings.Join(elem, " <- "))
		if err != nil {
			return err
		}
	}

	if info.Inner != nil {
		err := internal.WriteString(w, "\nCaused by:\n")
		if err != nil {
			return err
		}

		err = WriteError(w, info.Inner)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteError displays the error in full. It prints the error message and, if any,
// optional information.
//
// Parameters:
//   - w: The writer to write to.
//   - to_display: The error to display.
//
// Returns:
//   - error: If the error could not be displayed in full.
func WriteError(w io.Writer, to_display error) error {
	if to_display == nil {
		return nil
	}

	err := internal.WriteString(w, to_display.Error())
	if err != nil {
		return err
	}

	info, ok := to_display.(InfoWriter)
	if !ok {
		return nil
	}

	err = info.WriteInfo(w)
	return err
}

// Panic is like DisplayError but panics afterwards.
//
// Parameters:
//   - w: The writer to write to.
//   - to_display: The error to display.
func Panic(w io.Writer, to_display error) {
	if to_display == nil {
		return
	}

	info, ok := to_display.(InfoWriter)
	if ok {
		err := info.WriteInfo(w)
		if err != nil {
			panic(err)
		}
	}

	panic(to_display)
}

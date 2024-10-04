package errkit

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/PlayerR9/go-verify/test"
)

// TestWriteError tests the WriteError function.
func TestWriteError(t *testing.T) {
	err := errors.New("test error")

	var buff bytes.Buffer

	err = WriteError(&buff, err)
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}

	got := buff.String()
	if got != "test error" {
		t.Errorf("want %q, got %q", "test error", got)
	}
}

// TestPanic tests the Panic function.
func TestPanic(t *testing.T) {
	const (
		Expected string = "test error"
	)

	err := test.CheckPanic(Expected, func() {
		err := errors.New(Expected)

		Panic(io.Discard, err)
	})
	if err != nil {
		t.Error(err)
	}
}

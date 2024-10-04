package errkit

import (
	"bytes"
	"errors"
	"testing"
)

// TestWriteString tests the WriteString function.
func TestWriteString(t *testing.T) {
	const (
		TestData string = "This is a data used for testing purposes"
	)

	var buff bytes.Buffer

	err := WriteString(&buff, TestData)
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}

	got := buff.String()

	if got != TestData {
		t.Errorf("want %q, got %q", TestData, got)
	}
}

// TestDisplayError tests the DisplayError function.
func TestDisplayError(t *testing.T) {
	err := errors.New("test error")

	var buff bytes.Buffer

	err = DisplayError(&buff, err)
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}

	got := buff.String()
	if got != "test error" {
		t.Errorf("want %q, got %q", "test error", got)
	}
}

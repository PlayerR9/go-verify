package internal

import (
	"bytes"
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

// TestWriteStringf tests the WriteStringf function.
func TestWriteStringf(t *testing.T) {
	const (
		TestData string = "This is a data used for testing purposes"
	)

	var buff bytes.Buffer

	err := WriteStringf(&buff, TestData)
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}

	got := buff.String()

	if got != TestData {
		t.Errorf("want %q, got %q", TestData, got)
	}
}

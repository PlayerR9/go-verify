package test

import (
	"errors"
	"testing"
)

// TestTry tests the Try function.
func TestTry(t *testing.T) {
	var (
		Expected error = errors.New("something went wrong")
	)

	err := Try(func() {
		panic(Expected)
	})

	err = CheckErr(Expected.Error(), err)
	if err != nil {
		t.Error(err)
	}
}

// TestEqualsErr tests the EqualsErr function.
func TestEqualsErr(t *testing.T) {
	err1 := errors.New("this is an error")
	err2 := errors.New("this is an error")
	err3 := errors.New("this is another error")

	ok := EqualsErr(err1, err1)
	if !ok {
		t.Error("want true, got false")
	}

	ok = EqualsErr(err1, err2)
	if !ok {
		t.Error("want true, got false")
	}

	ok = EqualsErr(err1, err3)
	if ok {
		t.Error("want false, got true")
	}
}

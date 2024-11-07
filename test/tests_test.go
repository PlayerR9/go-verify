package test

import (
	"testing"

	"github.com/PlayerR9/go-verify/common"
)

func TestNewTests(t *testing.T) {
	type mock_args struct {
	}

	tests := NewTests[mock_args](nil)

	if tests.make_test(mock_args{}) == nil {
		t.Errorf("want UnimplementedTest, got nil")
	}

	if len(tests.tests) != 0 {
		t.Errorf("want 0, got %d", len(tests.tests))
	}

	var other_tests *Tests[mock_args]

	err := other_tests.AddTest("test", mock_args{})
	if err == nil {
		t.Errorf("want error, got nil")
	} else if err != common.ErrNilReceiver {
		t.Errorf("want ErrNilReceiver, got %v", err)
	}
}

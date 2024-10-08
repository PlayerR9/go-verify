package test

import (
	"testing"
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

	ok := other_tests.AddTest("test", mock_args{})
	if ok {
		t.Errorf("want false, got true")
	}
}

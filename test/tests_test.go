package test

import (
	"testing"

	"github.com/PlayerR9/go-verify/common"
)

// TestNewTestsSuccess tests the NewTests function when it succeeds.
func TestNewTestsSuccess(t *testing.T) {
	const (
		Want uint = 0
	)

	type mock_args struct{}

	tests := NewTests[mock_args](nil)

	fn := tests.make_test(mock_args{})
	if fn == nil {
		err := common.NewErrTestFailed("non-nil", "nil")
		t.Error(err)

		return
	}

	test_count := tests.GetTestsCount()
	if test_count == Want {
		return
	}

	common.FAIL.WrongInt(t, int(Want), int(test_count))
}

// TestNewTestsFail tests the NewTests function when it fails.
func TestNewTestsFail(t *testing.T) {
	type mock_args struct {
	}

	var tests *Tests[mock_args]

	err := tests.AddTest("test", mock_args{})
	if err == common.ErrNilReceiver {
		return
	}

	common.FAIL.WrongError(t, common.ErrNilReceiver.Error(), err)
}

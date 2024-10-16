package assert

import (
	"errors"
	"fmt"
)

type ErrPanic struct {
	Value any
}

func (e ErrPanic) Error() string {
	return fmt.Sprintf("%v", e.Value)
}

func NewErrPanic(value any) error {
	return &ErrPanic{
		Value: value,
	}
}

func try(err *error, fn func()) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch r := r.(type) {
		case string:
			*err = errors.New(r)
		case error:
			*err = r
		default:
			*err = NewErrPanic(r)
		}
	}()

	fn()
}

func Try(fn func()) error {
	if fn == nil {
		return nil
	}

	var err error

	try(&err, fn)

	return err
}

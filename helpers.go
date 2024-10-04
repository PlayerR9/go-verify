package errs

import (
	"reflect"
	"strings"

	"github.com/PlayerR9/go-verify/assert"
)

// OrElse is a helper function that either writes the first string
// to the builder if it is not empty, or the second string if it
// is empty.
//
// Parameters:
//   - b: The builder to write to.
//   - first: The first string to write.
//   - second: The second string to write.
//
// Does nothing if the builder is nil.
func OrElse(b *strings.Builder, first, second string) {
	if b == nil {
		return
	}

	if first == "" {
		b.WriteString(second)
	} else {
		b.WriteString(first)
	}
}

// StringOfType is a helper function that either writes the string
// representation of the value to the builder if it is not nil, or
// the string "nil" if it is nil.
//
// Parameters:
//   - b: The builder to write to.
//   - v: The value to write.
//
// Does nothing if the builder is nil.
func StringOfType(b *strings.Builder, v any) {
	if b == nil {
		return
	}

	if v == nil {
		b.WriteString("nil")
	}

	type_of := reflect.TypeOf(v)
	assert.Cond(type_of != nil, "type_of = nil")

	b.WriteString(type_of.String())
}

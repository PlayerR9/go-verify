package assert

// Validater is an interface that defines a method that checks the struct's
// inner state/integrity.
type Validater interface {
	// Validate validates the struct. Each implementation of Validater should
	// describe in the comments the conditions under which the struct is valid.
	//
	// Returns:
	//   - error: An error if the struct is invalid. Nil otherwise.
	//
	// NOTES: This method should not modify the struct's state. For that, use
	// Fixer instead.
	Validate() error
}

// Validate asserts whether the struct's inner state is valid. If not, it
// panics with an ErrValidationFailed error.
//
// Parameters:
//   - v: The struct to validate.
//   - name: The name of the struct. If empty, the name "struct" is used.
//   - allow_nil: Whether to allow the struct to be nil.
//
// Example:
//
//	type MyStruct struct {
//		Name string
//	}
//
//	func (ms MyStruct) Validate() error {
//		if ms.Name == "" {
//			return errors.New("name cannot be empty")
//		}
//
//		return nil
//	}
//
//	ms := &MyStruct{
//		Name: "",
//	}
//
//	Validate(ms, "ms", false) // Panics: (Validate Failed) ms = name cannot be empty
func Validate(v Validater, name string, allow_nil bool) {
	if v == nil && !allow_nil {
		panic(NewErrValidateFailed(name, nil))
	} else if v == nil {
		return
	}

	err := v.Validate()
	if err != nil {
		panic(NewErrValidateFailed(name, err))
	}
}

// Fixer is an interface that defines a method that validates and tries to bring the
// struct into the closest valid state.
type Fixer interface {
	// Fix fixes the struct. Each implementation of Fixer should describe in the
	// comments the conditions under which the struct is valid.
	//
	// Returns:
	//   - error: An error if the struct could not be fixed. Nil otherwise.
	Fix() error
}

// Fix fixes the struct to the closest valid state. Panics with an ErrValidationFailed
// error if the struct could not be fixed.
//
// Parameters:
//   - v: The struct to fix.
//   - name: The name of the struct. If empty, the name "struct" is used.
//   - allow_nil: Whether to allow the struct to be nil.
//
// Example:
//
//	type MyStruct struct {
//		Name string
//	}
//
//	func (ms *MyStruct) Fix() error {
//		if ms == nil {
//			return nil
//		}
//
//		if ms.Name == "" {
//			return errors.New("name cannot be empty")
//		}
//
//		return nil
//	}
//
//	ms := &MyStruct{
//		Name: "",
//	}
//
//	Fix(ms, "ms", false) // Panics: (Fix Failed) ms = name cannot be empty
func Fix(v Fixer, name string, allow_nil bool) {
	if v == nil && !allow_nil {
		panic(NewErrFixFailed(name, nil))
	} else if v == nil {
		return
	}

	err := v.Fix()
	if err != nil {
		panic(NewErrFixFailed(name, err))
	}
}

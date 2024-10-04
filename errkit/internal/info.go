package internal

import (
	"strings"
	"time"
)

// Info is a struct that holds information about an error.
type Info struct {
	// Context is a map of key-value pairs that can be used to store
	// additional information about the error.
	Context map[string]any

	// Suggestions is a list of suggestions that can be used to
	// solve the error.
	Suggestions []string

	// Frames is a list of stack frames that can be used to
	// determine where the error occurred.
	Frames []string

	// Inner is an optional field that can be used to store
	// the information about an error that is nested within
	// the current error.
	Inner error

	// Timestamp is the time at which the error was created.
	// It is used to determine the age of the error.
	Timestamp time.Time
}

// NewInfo creates a new Info struct with the current time.
//
// Returns:
//   - *Info: The new Info struct. Never returns nil.
func NewInfo() *Info {
	return &Info{
		Context:   make(map[string]any),
		Timestamp: time.Now(),
	}
}

// Copy creates a shallow copy of the info.
//
// Returns:
//   - *Info: The new Info struct. Never returns nil.
func (info Info) Copy() *Info {
	suggestions := make([]string, len(info.Suggestions))
	copy(suggestions, info.Suggestions)

	var context map[string]any

	if info.Context == nil {
		context = make(map[string]any)
	} else {
		context = make(map[string]any, len(info.Context))

		for key, value := range info.Context {
			context[key] = value
		}
	}

	var stack_trace []string

	if info.Frames == nil {
		stack_trace = make([]string, 0)
	} else {
		stack_trace = make([]string, len(info.Frames))
		copy(stack_trace, info.Frames)
	}

	return &Info{
		Suggestions: suggestions,
		Timestamp:   info.Timestamp,
		Context:     context,
		Frames:      stack_trace,
		Inner:       info.Inner,
	}
}

// AddSuggestion adds a suggestion to the error. Does nothing
// if the receiver is nil or the sentence is empty.
//
// Parameters:
//   - sentences: The sentences of the suggestion.
//
// Sentences are joined by a space.
func (info *Info) AddSuggestion(sentences ...string) {
	if info == nil || len(sentences) == 0 {
		return
	}

	info.Suggestions = append(info.Suggestions, strings.Join(sentences, " "))
}

// AddFrame appends a frame to the error. Does nothing
// if the receiver is nil.
//
// Parameters:
//   - frame: The frame of the error.
func (info *Info) AddFrame(frame string) {
	if info == nil {
		return
	}

	info.Frames = append(info.Frames, frame)
}

// SetInner sets the inner error. Does nothing if the receiver is nil.
//
// Parameters:
//   - inner: The inner error.
func (info *Info) SetInner(inner error) {
	if info == nil {
		return
	}

	info.Inner = inner
}

// Add adds a key-value pair to the set. Overwrites any
// existing value if the key already exists.
//
// Parameters:
//   - key: The key of the pair.
//   - value: The value of the pair.
//
// Returns:
//   - bool: True if the key was added, false otherwise.
func (info *Info) AddContext(key string, value any) bool {
	if info == nil {
		return false
	}

	info.Context[key] = value

	return true
}

// Get gets a key from the set.
//
// Returns:
//   - any: The value of the key, or the zero value if the key does not exist.
//   - bool: True if the key exists, false otherwise.
func (info Info) Get(key string) (any, bool) {
	val, ok := info.Context[key]
	return val, ok
}

// Has checks if the set has a key.
//
// Returns:
//   - bool: True if the key exists, false otherwise.
func (info Info) Has(key string) bool {
	_, ok := info.Context[key]
	return ok
}

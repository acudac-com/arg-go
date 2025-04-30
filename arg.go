package arg

import "fmt"

// The interface that all arguments must satisfy
type arg interface {
	Errors() []string
	Valid() bool
	Invalid() bool
}

// Returns true if any of the given arguments are invalid.
func Invalid(args ...arg) bool {
	for _, arg := range args {
		if len(arg.Errors()) > 0 {
			return true
		}
	}
	return false
}

// Returns true if all of the given arguments are valid.
func Valid(args ...arg) bool {
	for _, arg := range args {
		if len(arg.Errors()) > 0 {
			return false
		}
	}
	return true
}

// A generic argument.
type Arg[T any] struct {
	Value  T
	errors []string
}

// Returns a new instance of an argument with the type of the given value.
// This argument satisfies the arg interface needed for the Invalid method.
// You can embed this argument in any other struct to add custom methods to it.
func New[T any](value T) *Arg[T] {
	return &Arg[T]{value, []string{}}
}

// Returns the list of errors that this argument has if any.
func (a *Arg[T]) Errors() []string {
	return a.errors
}

// Returns true if no errors exist on the arg.
func (a *Arg[T]) Valid() bool {
	return len(a.errors) == 0
}

// Returns true if any errors exist on the arg.
func (a *Arg[T]) Invalid() bool {
	return len(a.errors) > 0
}

// Removes all errors
func (a *Arg[T]) ClearErrors() *Arg[T] {
	a.errors = []string{}
	return a
}

// Adds a new custom error
func (a *Arg[T]) AddError(msg string, args ...any) *Arg[T] {
	a.errors = append(a.errors, fmt.Sprintf(msg, args...))
	return a
}

// Changes the value to the specified value if condition is true.
func (a *Arg[T]) FallbackIf(fallback T, condition bool) *Arg[T] {
	if condition {
		a.Value = fallback
	}
	return a
}

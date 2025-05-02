package arg

import (
	"errors"
	"fmt"
	"strings"
)

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

// Returns the first error
func FirstError(args ...arg) error {
	for _, arg := range args {
		if len(arg.Errors()) > 0 {
			return errors.New(arg.Errors()[0])
		}
	}
	return nil
}

// Returns all errors as a new error
func AllErrors(args ...arg) error {
	var errs []string
	for _, arg := range args {
		errs = append(errs, arg.Errors()...)
	}
	return errors.New(strings.Join(errs, "; "))
}

// A generic argument.
type Arg[T any] struct {
	value  *T
	errors []string
}

// Returns a new instance of an argument with the type of the given value.
// This argument satisfies the arg interface needed for the Invalid method.
// You can embed this argument in any other struct to add custom methods to it.
func New[T any](value *T) *Arg[T] {
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

// Adds a new error if the condition is met
func (a *Arg[T]) AddErrorIf(condition bool, msg string, args ...any) *Arg[T] {
	a.errors = append(a.errors, fmt.Sprintf(msg, args...))
	return a
}

// Changes the value to the specified value if condition is true.
func (a *Arg[T]) FallbackIf(fallback T, condition bool) *Arg[T] {
	if condition {
		*a.value = fallback
	}
	return a
}

// Changes the value to the specified value if its nil.
func (a *Arg[T]) FallbackIfNil(fallback T) *Arg[T] {
	if a.value == nil {
		*a.value = fallback
	}
	return a
}

// A comparable arg
type ComparableArg[T comparable] struct {
	*Arg[T]
}

// Returns a comparable argument instance for the provided value.
func Comparable[T comparable](value *T) *ComparableArg[T] {
	custom := New(value)
	return &ComparableArg[T]{custom}
}

// Returns a comparable argument instance for the provided value.
func C[T comparable](value *T) *ComparableArg[T] {
	return Comparable(value)
}

// Sets the value to the specified value if its empty (e.g. 0 for number or "" for string)
func (a *ComparableArg[T]) Default(fallback T) *ComparableArg[T] {
	var zero T
	if a.value == nil || *a.value == zero {
		*a.value = fallback
	}
	return a
}

// Adds an error if the value is zero (e.g. 0 for number or "" for string)
func (a *ComparableArg[T]) Populated() *ComparableArg[T] {
	var zero T
	if a.value == nil || *a.value == zero {
		a.AddError("must be populated")
	}
	return a
}

// Adds an error if the value is not zero (e.g. 0 for number or "" for string)
func (a *ComparableArg[T]) Empty() *ComparableArg[T] {
	var zero T
	if a.value != nil && *a.value != zero {
		a.AddError("must be empty")
	}
	return a
}

// Adds an error if the value is not equal to one of the specified values.
func (a *ComparableArg[T]) Is(values ...T) *ComparableArg[T] {
	for _, val := range values {
		if val == *a.value {
			return a
		}
	}
	if len(values) == 1 {
		a.AddError("must be %v", values[0])
	} else if len(values) > 1 {
		a.AddError("must be one of %v", values)
	}
	return a
}

// Adds an error if the value is equal to one of the specified values.
func (a *ComparableArg[T]) IsNot(values ...T) *ComparableArg[T] {
	for _, val := range values {
		if val == *a.value {
			a.AddError("%v not allowed", values)
		}
	}
	return a
}

package arg

import "fmt"

type Arg interface {
	Errors() []string
	Valid() bool
	Invalid() bool
}

func Invalid(args ...Arg) bool {
	for _, arg := range args {
		if len(arg.Errors()) > 0 {
			return true
		}
	}
	return false
}

// An argument of a custom type like a struct.
type customArg[T any] struct {
	Value  T
	errors []string
}

type CustomArg[T any] struct {
	*customArg[T]
}

// Returns a new instance of an argument with a custom type like a struct.
func Custom[T any](value T) *CustomArg[T] {
	return &CustomArg[T]{&customArg[T]{value, []string{}}}
}

// Returns a new instance of an argument with a custom type like a struct.
func C[T any](value T) *CustomArg[T] {
	return Custom(value)
}

func (a *customArg[T]) Errors() []string {
	return a.errors
}

func (a *customArg[T]) Valid() bool {
	return len(a.errors) == 0
}

func (a *customArg[T]) Invalid() bool {
	return len(a.errors) > 0
}

func (a *customArg[T]) ClearErrors() *customArg[T] {
	a.errors = []string{}
	return a
}

func (a *customArg[T]) AddError(msg string, args ...any) *customArg[T] {
	a.errors = append(a.errors, fmt.Sprintf(msg, args...))
	return a
}

func (a *customArg[T]) FallbackIf(fallbackValue T, condition bool) *customArg[T] {
	if condition {
		a.Value = fallbackValue
	}
	return a
}

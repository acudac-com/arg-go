package arg

import "slices"

// The type of any number
type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// A generic number argument
type NumberArg[T number] struct {
	*Arg[T]
}

// Returns a number argument instance for the provided value.
func Number[T number](value T) *NumberArg[T] {
	return &NumberArg[T]{}
}

// Returns a number argument instance for the provided value.
func N[T number](value T) *NumberArg[T] {
	return &NumberArg[T]{}
}

// Adds an error if the number is zero.
func (a *NumberArg[T]) Populated() *NumberArg[T] {
	var zero T
	if a.Value == zero {
		a.AddError("must be populated")
	}
	return a
}

// Adds an error if the number is not zero.
func (a *NumberArg[T]) Empty() *NumberArg[T] {
	var zero T
	if a.Value != zero {
		a.AddError("must be empty")
	}
	return a
}

// Adds an error if the number is not equal to one of the specified values.
func (a *NumberArg[T]) Is(values ...T) *NumberArg[T] {
	if slices.Contains(values, a.Value) {
		return a
	}
	if len(values) == 1 {
		a.AddError("must be %v", values[0])
	} else if len(values) > 1 {
		a.AddError("must be one of %v", values)
	}
	return a
}

// Adds an error if the number equals one of the specified values.
func (a *NumberArg[T]) IsNot(values ...T) *NumberArg[T] {
	if slices.Contains(values, a.Value) {
		a.AddError("must not be one of %v", values)
		return a
	}
	return a
}

// Adds an error if the number is not less than the specified max.
func (a *NumberArg[T]) Lt(max T) *NumberArg[T] {
	if a.Value >= max {
		a.AddError("must be less than %v", max)
	}
	return a
}

// Adds an error if the number is not less than or equal to the specified max.
func (a *NumberArg[T]) Lte(max T) *NumberArg[T] {
	if a.Value > max {
		a.AddError("must be less than or equal to %v", max)
	}
	return a
}

// Adds an error if the number is not greater than the specified min.
func (a *NumberArg[T]) Gt(min T) *NumberArg[T] {
	if a.Value <= min {
		a.AddError("must be greater than %v", min)
	}
	return a
}

// Adds an error if the number is not greater than or equal to the specified min.
func (a *NumberArg[T]) Gte(min T) *NumberArg[T] {
	if a.Value < min {
		a.AddError("must be greater than or equal to %v", min)
	}
	return a
}

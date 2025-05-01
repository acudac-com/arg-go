package arg

// The type of any number
type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// A generic number argument
type NumberArg[T number] struct {
	*ComparableArg[T]
}

// Returns a number argument instance for the provided value.
func Number[T number](value *T) *NumberArg[T] {
	return &NumberArg[T]{C(value)}
}

// Returns a number argument instance for the provided value.
func N[T number](value *T) *NumberArg[T] {
	return &NumberArg[T]{C(value)}
}

// Adds an error if the number is not less than the specified max.
func (a *NumberArg[T]) Lt(max T) *NumberArg[T] {
	if *a.value >= max {
		a.AddError("must be less than %v", max)
	}
	return a
}

// Adds an error if the number is not less than or equal to the specified max.
func (a *NumberArg[T]) Lte(max T) *NumberArg[T] {
	if *a.value > max {
		a.AddError("must be less than or equal to %v", max)
	}
	return a
}

// Adds an error if the number is not greater than the specified min.
func (a *NumberArg[T]) Gt(min T) *NumberArg[T] {
	if *a.value <= min {
		a.AddError("must be greater than %v", min)
	}
	return a
}

// Adds an error if the number is not greater than or equal to the specified min.
func (a *NumberArg[T]) Gte(min T) *NumberArg[T] {
	if *a.value < min {
		a.AddError("must be greater than or equal to %v", min)
	}
	return a
}

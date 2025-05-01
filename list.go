package arg

import "slices"

type ListArg[T any] struct {
	*Arg[[]T]
}

// Returns a T list argument instance for the provided value.
func List[T any](value []T) *ListArg[T] {
	custom := New(&value)
	return &ListArg[T]{custom}
}

// Returns a T list argument instance for the provided value.
func SL[T any](value []T) *ListArg[T] {
	custom := New(&value)
	return &ListArg[T]{custom}
}

// Sets the value of the list if its empty.
func (a *ListArg[T]) Default(fallback []T) *ListArg[T] {
	if len(*a.value) == 0 {
		*a.value = fallback
	}
	return a
}

// Adds an error if the list is empty
func (a *ListArg[T]) Populated() *ListArg[T] {
	if len(*a.value) == 0 {
		a.AddError("must be populated")
	}
	return a
}

// Adds an error if the list is populated.
func (a *ListArg[T]) Empty() *ListArg[T] {
	if len(*a.value) != 0 {
		a.AddError("must be empty")
	}
	return a
}

// Adds an error if the list's length is not equal to the specified value.
func (a *ListArg[T]) LenEqs(length int) *ListArg[T] {
	if len(*a.value) != length {
		a.AddError("must contain %d values", length)
	}
	return a
}

// Adds an error if the list's length is not greater than the specified value.
func (a *ListArg[T]) LenGt(length int) *ListArg[T] {
	if len(*a.value) <= length {
		a.AddError("must contain more than %d values", length)
	}
	return a
}

// Adds an error if the list's length is not greater than or equal to the specified value.
func (a *ListArg[T]) LenGte(length int) *ListArg[T] {
	if len(*a.value) < length {
		a.AddError("must contain at least %d values", length)
	}
	return a
}

// Adds an error if the list's length is not less than the specified value.
func (a *ListArg[T]) LenLt(length int) *ListArg[T] {
	if len(*a.value) >= length {
		a.AddError("must contain less than %d values", length)
	}
	return a
}

// Adds an error if the list's length is not less than or equal to the specified value.
func (a *ListArg[T]) LenLte(length int) *ListArg[T] {
	if len(*a.value) > length {
		a.AddError("must contain at most %d values", length)
	}
	return a
}

type ComparableListArg[T comparable] struct {
	*ListArg[T]
}

// Returns a comparable T list argument instance for the provided value.
func ComparableList[T comparable](value []T) *ComparableListArg[T] {
	list := List(value)
	return &ComparableListArg[T]{list}
}

// Returns a comparable T list argument instance for the provided value.
func CL[T comparable](value []T) *ComparableListArg[T] {
	return ComparableList(value)
}

// Adds an error if the list does not include all of the specified values
func (a *ComparableListArg[T]) Includes(values ...T) *ComparableListArg[T] {
	for _, val := range values {
		if !slices.Contains(*a.value, val) {
			a.AddError("must include %v", val)
		}
	}
	return a
}

// Sets each value to the specified value if its empty (e.g 0 for number or "" for string)
func (a *ComparableListArg[T]) EachDefault(fallback T) *ComparableListArg[T] {
	var zero T
	for i := range *a.value {
		if (*a.value)[i] == zero {
			(*a.value)[i] = fallback
		}
	}
	return a
}

// Adds an error if any of the values are zero (e.g 0 for number or "" for string)
func (a *ComparableListArg[T]) EachPopulated() *ComparableListArg[T] {
	var zero T
	for _, val := range *a.value {
		if val == zero {
			a.AddError("each value must be populated")
		}
	}
	return a
}

// Adds an error if any of the values are not zero (e.g 0 for number or "" for string)
func (a *ComparableListArg[T]) EachEmpty() *ComparableListArg[T] {
	var zero T
	for _, val := range *a.value {
		if val != zero {
			a.AddError("each value must be empty")
		}
	}
	return a
}

// Adds an error if any of the values are not equal to one of the specified values.
func (a *ComparableListArg[T]) EachIs(values ...T) *ComparableListArg[T] {
	for _, val := range *a.value {
		if !slices.Contains(values, val) {
			a.AddError("each value must be one of %v", values)
		}
	}
	return a
}

// Adds an error if any of the values are equal to one of the specified values.
func (a *ComparableListArg[T]) EachIsNot(values ...T) *ComparableListArg[T] {
	for _, val := range *a.value {
		if slices.Contains(values, val) {
			a.AddError("%v not allowed", val)
		}
	}
	return a
}

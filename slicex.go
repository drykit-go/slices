package slicex

type (
	// FilterFunc is a function that returns a boolean depending on v.
	FilterFunc[T any] func(v T) bool

	// MapFunc is a function that transforms v into a value of any type.
	MapFunc[T, U any] func(v T) U

	// ReduceFunc is a reducer function.
	ReduceFunc[U, T any] func(U, T) U
)

// Map applies f on each element of src and returns the resulting slice.
// The output is guaranteed to be the same length as src.
// src remains unaltered.
func Map[T, U any](src []T, f MapFunc[T, U]) []U {
	out := make([]U, len(src))
	for i, v := range src {
		out[i] = f(v)
	}
	return out
}

// Filter filters out elements of src for which f(element) returns false
// and returns the resulting slice.
// The output length is inferior or equal to src's length.
// src remains unaltered.
func Filter[T any](src []T, f FilterFunc[T]) []T {
	out := []T{}
	for _, v := range src {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce applies reducer f to src starting from ini and returns
// the accumulated value.
func Reduce[Elem, Accumulator any](
	src []Elem,
	f ReduceFunc[Accumulator, Elem],
	ini Accumulator,
) Accumulator {
	out := ini
	for _, v := range src {
		out = f(out, v)
	}
	return out
}

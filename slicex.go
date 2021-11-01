package slicex

// MapFunc is a function that transforms v into a value of any type.
type MapFunc[T, U any] func(v T) U

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

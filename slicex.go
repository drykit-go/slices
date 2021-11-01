package slicex

type (
	// FilterFunc is a function that returns a boolean value
	// that depends on the current element.
	FilterFunc[Elem any] func(Elem) bool

	// MapFunc is a function that transforms the current element into
	// a new value of any type.
	MapFunc[Elem, NewElem any] func(Elem) NewElem

	// ReduceFunc is a reducer function: it returns an accumulated value
	// that depends on the previously accumulated value and the current element.
	ReduceFunc[Accumulator, Elem any] func(Accumulator, Elem) Accumulator
)

// Map applies f on each element of src and returns the resulting slice.
// The output is guaranteed to be the same length as src.
// src remains unaltered.
func Map[Elem, NewElem any](src []Elem, f MapFunc[Elem, NewElem]) []NewElem {
	out := make([]NewElem, len(src))
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

// Apply iterates over src and calls f(currentIndex, currentElement)
// each iteration until the end is reached.
func Apply[Elem any](src []Elem, f func(i int, v Elem)) {
	for i, v := range src {
		f(i, v)
	}
}

// ApplyUntil iterates over src and calls f(currentIndex, currentElement)
// each iteration until it returns false or the end is reached.
func ApplyUntil[Elem any](src []Elem, f func(i int, v Elem) bool) {
	for i, v := range src {
		if !f(i, v) {
			return
		}
	}
}

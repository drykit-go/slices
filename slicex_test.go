package slicex_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/drykit-go/slicex"
)

func TestMap(t *testing.T) {
	ints := []int{0, 5, -1}
	toString := func(n int) string { return fmt.Sprint(n) }
	exp := []string{"0", "5", "-1"}
	got := slicex.Map(ints, toString)
	assertEqualSlices(t, got, exp)
}

func TestFilter(t *testing.T) {
	ints := []int{0, 5, -1}
	isEven := func(n int) bool { return n&1 == 0 }
	exp := []int{0}
	got := slicex.Filter(ints, isEven)
	assertEqualSlices(t, got, exp)
}

func assertEqualSlices[T any](t *testing.T, a, b []T) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("exp equal slices, got:\na == %v\nb == %v", a, b)
	}
}

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

func TestReduce(t *testing.T) {
	type bill struct {
		label  string
		amount float64
	}
	bills := []bill{
		{label: "a", amount: 3.14},
		{label: "b", amount: 39},
		{label: "c", amount: -.14},
	}
	calculateSum := func(sum float64, cur bill) float64 {
		return sum + cur.amount
	}
	exp := 42.
	if got := slicex.Reduce(bills, calculateSum, 0); got != exp {
		t.Errorf("exp %v\ngot %v", exp, got)
	}
}

func TestApply(t *testing.T) {
	s := make([]bool, 5)
	n := 0
	f := func(i int, v bool) {
		if !v {
			n = i + 1
		}
	}
	exp := 5
	slicex.Apply(s, f)
	if n != exp {
		t.Errorf("exp %v\ngot %v", exp, n)
	}
}

// Helpers

func assertEqualSlices[T any](t *testing.T, a, b []T) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("exp equal slices, got:\na == %v\nb == %v", a, b)
	}
}

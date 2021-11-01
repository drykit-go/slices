package slicex_test

import (
	"fmt"
	"reflect"
	"sort"
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

func TestApplyUntil(t *testing.T) {
	newSlice := func() []int {
		return []int{0, 1, 2, 3}
	}
	t.Run("stop at index", func(t *testing.T) {
		n := 0
		slicex.ApplyUntil(newSlice(), func(i, _ int) bool {
			n = i
			return i < 2
		})
		exp := 2
		if n != exp {
			t.Errorf("did not stop at expected index:\nexp %v\ngot %v", exp, n)
		}
	})
	t.Run("stop at value", func(t *testing.T) {
		n := 0
		slicex.ApplyUntil(newSlice(), func(_, v int) bool {
			n = v
			return v < 2
		})
		exp := 2
		if n != exp {
			t.Errorf("did not stop at expected value:\nexp %v\ngot %v", exp, n)
		}
	})
	t.Run("stop at end", func(t *testing.T) {
		n := 0
		slicex.ApplyUntil(newSlice(), func(i, _ int) bool {
			n = i
			return true
		})
		exp := 3
		if n != exp {
			t.Errorf("did stop before the end:\nexp index %v\ngot index %v", exp, n)
		}
	})
}

func TestKeysOf(t *testing.T) {
	m := map[string]interface{}{
		"a": 3,
		"b": "hi",
		"c": false,
	}
	keys := slicex.KeysOf(m)
	sort.Strings(keys)
	exp := []string{"a", "b", "c"}
	assertEqualSlices(t, keys, exp)
}

// Helpers

func assertEqualSlices[T any](t *testing.T, a, b []T) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("exp equal slices, got:\na == %v\nb == %v", a, b)
	}
}

package arraysandslices

import (
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("sum of one slice equals just the sum of that slice's elements", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5})
		want := []int{15}
		assertSlicesEqual(t, got, want)
	})
	t.Run("sum of 0 slices is empty slice", func(t *testing.T) {
		got := SumAll()
		want := []int{}
		assertSlicesEqual(t, got, want)
	})
	t.Run("sum of multiple slices is the sum of the sums of their elements", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2})
		want := []int{15, 3}
		assertSlicesEqual(t, got, want)
	})
}

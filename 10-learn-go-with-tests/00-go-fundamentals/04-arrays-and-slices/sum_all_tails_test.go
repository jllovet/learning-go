package arraysandslices

import "testing"

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	want := []int{14, 5}
	assertSlicesEqual(t, got, want)
}

func TestTail(t *testing.T) {
	t.Run("tail of a non-empty slice returns all but first element", func(t *testing.T) {
		got := tail([]int{1, 2, 3, 4, 5})
		want := []int{2, 3, 4, 5}
		assertSlicesEqual(t, got, want)
	})
	t.Run("tail of an empty slice returns empty slice", func(t *testing.T) {
		got := tail([]int{})
		want := []int{}
		assertSlicesEqual(t, got, want)
	})
}

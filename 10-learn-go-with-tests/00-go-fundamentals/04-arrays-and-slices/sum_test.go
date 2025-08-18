package arraysandslices

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSum() {
	n := []int{1, 2, 3, 4, 5}
	sum := Sum(n)
	fmt.Println(sum)
	// Output: 15
}

func TestSum(t *testing.T) {
	t.Run("sum of 1...5 is 15", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4, 5})
		want := 15
		assertCorrectMessage(t, got, want)
	})
	t.Run("sum of empty slice is 0", func(t *testing.T) {
		got := Sum([]int{})
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("sum of slice with one element is that element", func(t *testing.T) {
		got := Sum([]int{2})
		want := 2
		assertCorrectMessage(t, got, want)
	})
}
func TestRecursiveSum(t *testing.T) {
	t.Run("sum of 1...5 is 15", func(t *testing.T) {
		got := recursiveSum([]int{1, 2, 3, 4, 5})
		want := 15
		assertCorrectMessage(t, got, want)
	})
	t.Run("sum of empty slice is 0", func(t *testing.T) {
		got := recursiveSum([]int{})
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("sum of slice with one element is that element", func(t *testing.T) {
		got := recursiveSum([]int{2})
		want := 2
		assertCorrectMessage(t, got, want)
	})
}
func TestIterativeSum(t *testing.T) {
	t.Run("sum of 1...5 is 15", func(t *testing.T) {
		got := iterativeSum([]int{1, 2, 3, 4, 5})
		want := 15
		assertCorrectMessage(t, got, want)
	})
	t.Run("sum of empty slice is 0", func(t *testing.T) {
		got := iterativeSum([]int{})
		want := 0
		assertCorrectMessage(t, got, want)
	})
	t.Run("sum of slice with one element is that element", func(t *testing.T) {
		got := iterativeSum([]int{2})
		want := 2
		assertCorrectMessage(t, got, want)
	})
}

// Benchmarking two alternative implementations

func BenchmarkRecursiveSum(b *testing.B) {
	n := func(size int) []int {
		fmt.Printf("benchmarking recursiveSum with array of %d integers\n", size)
		var s []int
		for v := range size {
			s = append(s, v)
		}
		return s
	}(100000)
	for b.Loop() {
		recursiveSum(n)
	}
}

func BenchmarkIterativeSum(b *testing.B) {
	n := func(size int) []int {
		fmt.Printf("benchmarking iterativeSum with array of %d integers\n", size)
		var s []int
		for v := range size {
			s = append(s, v)
		}
		return s
	}(100000)
	for b.Loop() {
		iterativeSum(n)
	}
}

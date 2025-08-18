package arraysandslices

func recursiveSum(n []int) int {
	// Benchmark with 100000 ints, goos: darwin, goarch: arm64, cpu: Apple M2 Pro
	// BenchmarkRecursiveSum-10		1838		643367 ns/op		0 B/op		0 allocs/op
	var sumHelper func([]int, int) int
	sumHelper = func(s []int, idx int) int {
		switch idx {
		case len(s):
			return 0
		default:
			return s[idx] + sumHelper(s, idx+1)
		}
	}
	return sumHelper(n, 0)
}

func iterativeSum(n []int) int {
	// Benchmark with 100000 ints, goos: darwin, goarch: arm64, cpu: Apple M2 Pro
	// BenchmarkIterativeSum-10		31897		37547 ns/op			0 B/op		0 allocs/op
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}

func Sum(n []int) int {
	return iterativeSum(n)
}

func SumAll(ns ...[]int) []int {
	sums := make([]int, len(ns))
	for i, n := range ns {
		sums[i] = Sum(n)
	}
	return sums
}

func tail(n []int) []int {
	switch len(n) {
	case 0:
		return []int{}
	default:
		return n[1:]
	}
}

func SumAllTails(ns ...[]int) []int {
	tails := make([]int, len(ns))
	for i, n := range ns {
		tails[i] = Sum(tail(n))
	}
	return tails
}

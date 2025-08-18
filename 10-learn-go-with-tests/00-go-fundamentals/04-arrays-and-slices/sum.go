package main

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

# Profiling Before and After Refactoring to Use Goroutines with Channels

goos: darwin
goarch: arm64
pkg: github.com/jllovet/learning-go/10-learn-go-with-tests/00-go-fundamentals/10-concurrency
cpu: Apple M2 Pro
BenchmarkCheckWebsites-10    	       1	2090196292 ns/op	     384 B/op	       5 allocs/op
PASS
ok  	github.com/jllovet/learning-go/10-learn-go-with-tests/00-go-fundamentals/10-concurrency	2.291s


goos: darwin
goarch: arm64
pkg: github.com/jllovet/learning-go/10-learn-go-with-tests/00-go-fundamentals/10-concurrency
cpu: Apple M2 Pro
BenchmarkCheckWebsites-10    	      56	  21171318 ns/op	   20537 B/op	     224 allocs/op
PASS
ok  	github.com/jllovet/learning-go/10-learn-go-with-tests/00-go-fundamentals/10-concurrency	1.381s
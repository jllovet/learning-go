package iteration

import "strings"

func Repeat(s string, count int) string {
	// see strings.Repeat
	var repeated string
	for range count {
		repeated += s
	}
	return repeated
}

func StdLibRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

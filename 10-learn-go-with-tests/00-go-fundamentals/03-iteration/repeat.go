package iteration

import "strings"

func Repeat(s string, count int) string {
	var repeated string
	for range count {
		repeated += s
	}
	return repeated
}

func BuilderRepeat(s string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(s)
	}
	return repeated.String()
}

func StdLibRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

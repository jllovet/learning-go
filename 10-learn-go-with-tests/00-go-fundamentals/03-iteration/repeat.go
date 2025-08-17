package iteration

func Repeat(s string, count int) string {
	// see strings.Repeat
	var repeated string
	for range count {
		repeated += s
	}
	return repeated
}

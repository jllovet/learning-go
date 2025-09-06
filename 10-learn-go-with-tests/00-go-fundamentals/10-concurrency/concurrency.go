package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	type result struct {
		url   string
		valid bool
	}
	resultChannel := make(chan result)
	results := make(map[string]bool)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url: url, valid: wc(url)}
		}()
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.url] = r.valid
	}

	return results
}

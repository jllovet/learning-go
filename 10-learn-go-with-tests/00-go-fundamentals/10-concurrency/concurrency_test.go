package concurrency_test

import (
	"reflect"
	"testing"
	"time"

	concurrency "github.com/jllovet/learning-go/10-learn-go-with-tests/00-go-fundamentals/10-concurrency"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"https://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"https://google.com":         true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}
	got := concurrency.CheckWebsites(mockWebsiteChecker, urls)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range len(urls) {
		urls[i] = "a url"
	}

	for b.Loop() {
		concurrency.CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

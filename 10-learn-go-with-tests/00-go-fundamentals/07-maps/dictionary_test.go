package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Dictionary wrapper type", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		searchTerm := "test"
		got := dictionary.Search(searchTerm)
		want := "this is just a test"
		assertString(t, got, want)
	})

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	searchTerm := "test"
	got := Search(dictionary, searchTerm)
	want := "this is just a test"
	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

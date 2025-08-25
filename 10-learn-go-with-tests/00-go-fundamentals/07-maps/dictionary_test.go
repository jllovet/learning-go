package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	searchTerm := "test"
	got := Search(dictionary, searchTerm)
	want := "this is just a test"
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, searchTerm)
	}
}

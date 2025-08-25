package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		searchTerm := "test"
		got, _ := dictionary.Search(searchTerm)
		want := "this is just a test"
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		searchTerm := "unknown"
		_, err := dictionary.Search(searchTerm)
		if err == nil {
			t.Fatal("expected to get an error but didn't get one")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add a word", func(t *testing.T) {
		d := Dictionary{}
		d.Add("test", "this is just a test")
		want := "this is just a test"
		got, err := d.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertString(t, got, want)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

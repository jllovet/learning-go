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
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := d.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		err := d.Add(word, "new definition")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		d := Dictionary{word: definition}
		err := d.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, d, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition"
		d := Dictionary{}
		err := d.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		err := d.Delete(word)
		assertError(t, err, nil)
		_, err = d.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("non-existent word", func(t *testing.T) {
		word := "test"
		d := Dictionary{}
		err := d.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, definition)
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

package arraysandslices

import (
	"slices"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertSlicesEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

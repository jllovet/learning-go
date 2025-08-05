package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Carl")
	want := "Hello, Carl"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package smi

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assertCorrectMessage(t, got, want)
}

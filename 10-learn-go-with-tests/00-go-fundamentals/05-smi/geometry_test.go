package smi

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0
	assertCorrectMessage(t, got, want)
}

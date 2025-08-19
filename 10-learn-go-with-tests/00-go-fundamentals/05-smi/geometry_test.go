package smi

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{Length: 10.0, Width: 10.0}
	got := Perimeter(r)
	want := 40.0
	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	r := Rectangle{Length: 12.0, Width: 6.0}
	got := Area(r)
	want := 72.0
	assertCorrectMessage(t, got, want)
}

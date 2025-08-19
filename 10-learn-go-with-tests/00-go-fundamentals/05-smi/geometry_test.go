package smi

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{Length: 10.0, Width: 10.0}
	got := r.Perimeter()
	want := 40.0
	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{Length: 12.0, Width: 6.0}
		got := r.Area()
		want := 72.0
		assertCorrectMessage(t, got, want)
	})
	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		got := c.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

package smi

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{Length: 10.0, Width: 10.0}
	got := r.Perimeter()
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle 10x10", shape: Rectangle{Length: 10.0, Width: 10.0}, want: 100.0},
		{name: "rectangle 2x4", shape: Rectangle{Length: 2.0, Width: 4.0}, want: 8.0},
		{name: "circle r=4", shape: Circle{Radius: 4.0}, want: 16 * math.Pi},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			checkArea(t, tt.shape, tt.want)
		})
	}
}

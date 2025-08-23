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
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle 10x10", shape: Rectangle{Length: 10.0, Width: 10.0}, hasArea: 100.0},
		{name: "rectangle 2x4", shape: Rectangle{Length: 2.0, Width: 4.0}, hasArea: 8.0},
		{name: "circle r=4", shape: Circle{Radius: 4.0}, hasArea: 16 * math.Pi},
		{name: "triangle b=2 h=3", shape: Triangle{Base: 2.0, Height: 3.0}, hasArea: 3.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			checkArea(t, tt.shape, tt.hasArea)
		})
	}
}

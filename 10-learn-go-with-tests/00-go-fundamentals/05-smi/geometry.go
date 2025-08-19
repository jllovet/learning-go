package smi

type Rectangle struct {
	Length float64
	Width  float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Length + r.Width)
}

func Area(r Rectangle) float64 {
	return r.Length * r.Width
}

package shapes

type Shape interface {
	Area() float64
}

func ShapeArea(s Shape) float64 {
	return s.Area()
}

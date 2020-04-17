package shapes

type Triangle struct {
	base   float64
	height float64
}

func NewTriangle(b float64, h float64) *Triangle {
	return &Triangle{
		base:   b,
		height: h,
	}
}

func (t *Triangle) Area() float64 {
	return (t.base * t.height) / 2
}

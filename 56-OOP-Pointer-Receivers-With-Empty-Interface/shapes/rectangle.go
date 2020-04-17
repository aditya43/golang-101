package shapes

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{
		width:  w,
		height: h,
	}
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

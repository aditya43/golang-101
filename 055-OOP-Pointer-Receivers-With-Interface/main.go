package main

import (
	"fmt"

	"github.com/aditya43/golang/55-OOP-Pointer-Receivers-With-Interface/shapes"
)

func main() {
	r := shapes.NewRectangle(9, 6)
	t := shapes.NewTriangle(3, 6)

	fmt.Println("Rectangle Area", shapes.ShapeArea(r))
	fmt.Println("Trianlge Area", shapes.ShapeArea(t))
}

package main

import (
	"fmt"

	"github.com/aditya43/golang/55-OOP-Pointer-Receivers/shapes/shapeInterface"
)

func main() {
	r := shapeInterface.NewRectangle(9, 6)
	t := shapeInterface.NewRectangle(3, 6)

	fmt.Println("Rectangle Area", r.ShapeArea())
	fmt.Println("Trianlge Area", t.ShapeArea())
}

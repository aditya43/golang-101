package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aditya43/golang/55-OOP-Pointer-Receivers-With-Interface/shapes"
)

func main() {
	myRectangle := shapes.NewRectangle(4, 5)
	myTriangle := shapes.NewTriangle(2, 7)
	shapesSlice := []interface{}{myRectangle, myTriangle}

	for index, shape := range shapesSlice {
		fmt.Printf("Shape in index, %d, of the shapesSlice is a  %T\n", index, shape)
	}
	fmt.Println()

	aRandomShape := getRandomShape()

	fmt.Printf("The type of aRandomShape is %T\n", aRandomShape)

	switch t := aRandomShape.(type) {
	case *shapes.Rectangle:
		fmt.Println("I got back a rectangle with an area equal to ", t.Area())
	case *shapes.Triangle:
		fmt.Println("I got back a triangle with an area equal to ", t.Area())
	default:
		fmt.Println("I don't recognize what I got back!")
	}
}

// Returning empty interface
func getRandomShape() interface{} {
	var shape interface{}

	var randomShapesSlice []interface{} = make([]interface{}, 2)

	// Seed random number generator
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// Pick a random number, either 1 or 0
	randomNumber := r.Intn(2)
	fmt.Println("Random Number: ", randomNumber)

	// Let's make a new rectangle
	rectangle := shapes.NewRectangle(3, 6)

	// Let's make a new triangle
	triangle := shapes.NewTriangle(9, 18)

	// Let's store the shapes into a slice data structure
	randomShapesSlice[0] = rectangle
	randomShapesSlice[1] = triangle
	shape = randomShapesSlice[randomNumber]

	return shape
}

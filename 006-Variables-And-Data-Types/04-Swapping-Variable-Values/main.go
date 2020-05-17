package main

import "fmt"

func main() {
	firstName, lastName := "Aditya", "Hajare"
	fmt.Println("Before Swapping: ", firstName, lastName)

	firstName, lastName = lastName, firstName
	fmt.Println("After Swapping: ", firstName, lastName)
}

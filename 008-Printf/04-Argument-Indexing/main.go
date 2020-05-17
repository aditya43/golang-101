package main

import "fmt"

func main() {
	var (
		firstName = "Aditya"
		age       = 33
		isActive  = true
	)

	// Using Argument indexing in Printf. Argument indexing starts at 1
	fmt.Printf("\nName: %[3]v | Age: %[1]v | isActive: %[2]v\n", firstName, age, isActive)
}

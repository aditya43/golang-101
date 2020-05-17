package main

import "fmt"

func main() {
	var (
		firstName = "Aditya"
		age       = 33
		isActive  = true
	)

	// Using Println:
	fmt.Println("Name:", firstName, " | Age:", age, " | isActive:", isActive)

	// Using Printf
	fmt.Printf("\nName: %s | Age: %d | isActive: %t\n", firstName, age, isActive)

	// Using Printf - If we don't know the value type, use %v
	fmt.Printf("\nName: %v | Age: %v | isActive: %v\n", firstName, age, isActive)
}

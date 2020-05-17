package main

import "fmt"

func main() {
	msg := "नमस्ते आदित्य!"

	// Closure function
	func(m string) {
		fmt.Println(m)
	}(msg)

	fmt.Println("Inside main() function")
}

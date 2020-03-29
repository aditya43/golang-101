package main

import "fmt"

func main() {
	// without newline
	fmt.Println("AdityaHajare")

	// with newline:
	//   \n = escape sequence
	//   \  = escape character
	fmt.Println("Aditya\nHajare")

	// escape characters:
	//   \\ = \
	//   \" = "
	fmt.Println("Aditya\\n\"nHajare\"")
}

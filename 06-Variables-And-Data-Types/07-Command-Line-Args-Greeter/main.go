package main

import (
	"fmt"
	"os"
)

// I didn't include any error detection here.
//
// So, if we don't pass any command line argument, this program will fail.
// This is intentional.

func main() {
	// fmt.Printf("%#v\n", os.Args) // To see whats inside 'os.Args'

	name := os.Args[1]

	fmt.Println("नमस्ते", name)
}

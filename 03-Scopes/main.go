package main

// package scope
// file scope

import "fmt"

const ok = true

func main() { // block scope starts

	var namaste = "नमस्ते आदित्य"

	// namaste and ok are visible here
	fmt.Println(namaste, ok)

} // block scope ends

package main

import "fmt"

func main() {
	str := "नमस्ते आदित्य"

	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("% x\n", bytes)
}
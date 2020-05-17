package main

import "fmt"

func main() {
	floatVal := 123.456

	// Float Precisions
	fmt.Printf("Float Value: %f\n", floatVal)
	fmt.Printf("Float Value: %.0f\n", floatVal)
	fmt.Printf("Float Value: %.1f\n", floatVal)
	fmt.Printf("Float Value: %.2f\n", floatVal)
	fmt.Printf("Float Value: %.3f\n", floatVal)
	fmt.Printf("Float Value: %.4f\n", floatVal)
	fmt.Printf("Float Value: %.5f\n", floatVal)
	fmt.Printf("Float Value: %.6f\n", floatVal)
}

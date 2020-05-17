package main

import "fmt"

func main() {
	var (
		adiBool    bool
		adiInt     int
		adiFloat   float64
		adiStr     string
		adiPointer *string
	)

	fmt.Printf("\nadiBool: %T", adiBool)
	fmt.Printf("\nadiInt: %T", adiInt)
	fmt.Printf("\nadiFloat: %T", adiFloat)
	fmt.Printf("\nadiStr: %T", adiStr)
	fmt.Printf("\nadiPointer: %T\n", adiPointer)
}

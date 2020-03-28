package main

import "fmt"

func main() {
	var adiBool bool       // false
	var adiInt int         // 0
	var adiFloat float64   // 0
	var adiStr string      // ""
	var adiPointer *string // nil | 'nil' means it doesn't point to any memory location

	fmt.Println("Zero Values assigned to variables based on their types:")

	fmt.Println("Bool", adiBool)
	fmt.Println("Int", adiInt)
	fmt.Println("Float64", adiFloat)
	fmt.Println("String", adiStr)
	fmt.Println("Pointer", adiPointer)
}

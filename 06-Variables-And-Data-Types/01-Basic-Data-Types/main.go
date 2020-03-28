package main

import "fmt"

func main() {
	integers()
	floats()
	booleans()
	strings()
	hexadecimals()
}

// Print integer literals.
func integers() {
	fmt.Println(
		"Integer Literals: ",
		-200, -100, 0, 10, 20, 1000, // Ending comma is required since closing parenthesis is not on this line.
	)
}

// Print float literals.
func floats() {
	fmt.Println(
		"Float Literals: ",
		-100.5,
		-10.2,
		-5.300,
		-4.,
		2.0,
		20,
		30.100, // Ending comma is required since closing parenthesis is not on this line.
	)
}

// Print boolean literals.
func booleans() {
	fmt.Println(
		"Boolean Literals: ", true, false, // Ending comma is required since closing parenthesis is not on this line.
	)
}

// Print string literals.
func strings() {
	fmt.Println("String Literals: ", "", "Aditya", "नमस्ते आदित्य") // Ending comma is not required since closing parenthesis is on the same line.
}

// Print hexadecimal literals.
func hexadecimals() {
	fmt.Println("Hexadecimal Literals: ", 0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9) // 0 1 2 3 4 5 6 7 8 9

	fmt.Println("Hexadecimal Literals: ", 0xa, 0xb, 0xc, 0xd, 0xe, 0xf) // 10 11 12 13 14 15

	fmt.Println("Hexadecimal Literal: ", 0x11) // 17
	fmt.Println("Hexadecimal Literal: ", 0x19) // 25
	fmt.Println("Hexadecimal Literal: ", 0x32) // 50
	fmt.Println("Hexadecimal Literal: ", 0x64) // 100
}

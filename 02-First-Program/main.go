// package main is a special package
// it allows Go to create an executable file
package main

/*
This is a multi-line comment.
import keyword makes another package availablec for this .go "file".
import "fmt" lets you access fmt package's functionality here in this file.
*/
import "fmt"

// "func main" is special.
// Go has to know where to start
// func main() creates a starting point for Go
//
// When code is compiled, Go runtime will first run main() function
func main() {
	// after: import "fmt"
	// Println function of "fmt" package becomes available

	// To see 'Println' docs, type below command in shell:
	// go doc -src fmt Println

	// Println is just an exported function from "fmt" package

	// Exported = First Letter is uppercase
	fmt.Println("Hello Aditya")
	fmt.Println("नमस्ते आदित्य")

	// Go cannot call Println function by itself.
	// That's why you need to call it here.
	// It only calls `func main` automatically.

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

	// Go supports Unicode characters in string literals
	// And also in source-code: आदित्य!
	//
	// Because: Literal ~= Source Code
}

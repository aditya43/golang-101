package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// strings are made up of bytes

	// len function counts the bytes in a string value.
	//
	// This string literal contains unicode characters.
	//
	// And, unicode characters can be 1-4 bytes.
	// So, "आदित्य" is 18 bytes long and has 6 characters.
	name := "आदित्य"
	fmt.Printf("%q is %d bytes\n", name, len(name))

	// To get the actual characters (or runes) inside
	// a utf-8 encoded string value, you should do this:
	fmt.Printf("%q is %d characters\n",
		name, utf8.RuneCountInString(name))
}

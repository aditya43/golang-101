package printer

// NOTE: This file must exist under 'printer' directory since we have named this package 'printer'.

import "fmt"

func namaste() { // Will not be exported since func name doesn't start with upper case letter.
	fmt.Println("नमस्ते from unexported function namaste")
}

// Namaste func will be exported since it's name starts with upper case letter and it will be available throughout 'printer' package.
func Namaste() {
	fmt.Println("नमस्ते from exported function namaste")
}

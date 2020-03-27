package printer

import "fmt"

func namaste() { // Will not be exported since func name doesn't start with upper case letter.
	fmt.Println("नमस्ते from unexported function namaste")
}

func Namaste() { // Since func name starts with upper case letter, it will be exported and be available throughout 'printer' package .
	fmt.Println("नमस्ते from unexported function namaste")
}

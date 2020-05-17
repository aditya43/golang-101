package main

import "fmt"

func main() {
	speed := 100 // int
	force := 2.5 // float64

	// speed = speed * force // ERROR: invalid op

	// conversion can be a destructive operation
	speed = speed * int(force) // `force` loses its fractional part!

	fmt.Println(speed) // This will give wrong result (200) since 'force' variable lost its fractional part and multiplication happend as 100 x 2.
}

package main

import "fmt"

func main() {
	name, last := "आदित्य", "Hajare"

	// assignment operation using string concat
	name += " Doe"

	// equals to this:
	// name = name + " Doe"

	fmt.Println(name + " " + last)
}

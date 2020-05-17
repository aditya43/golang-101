package main

import "fmt"

func main() {
	// this one uses a raw string literal

	// can we see how readable it is?
	// compared to the previous one?

	json := `
{
	"Items": [{
		"Item": {
			"name": "Aditya Hajare"
		}
	}]
}`

	fmt.Println(json)
}

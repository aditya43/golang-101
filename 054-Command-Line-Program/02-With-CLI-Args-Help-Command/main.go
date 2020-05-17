// To execute this program:

// Method #1
// Run "go install" and then execute

// Method #2
// go run main.go
// Also run: "go run main.go --help"

package main

import (
	"flag"
	"fmt"
)

func main() {
	var userName string

	flag.StringVar(&userName, "Your Name", "Unknown User", "Enter your name")
	flag.Parse()

	fmt.Println("नमस्ते", userName)
}

package main

import (
	"fmt"
)

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("नमस्ते आदित्य!", i, source)
	}
}

func main() {
	go printGreetings("From Goroutine")
	printGreetings("From Main Function")
}

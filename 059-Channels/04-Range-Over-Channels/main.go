package main

import "fmt"

func main() {
	// Created buffered channel of strings with a capacity of 3
	// This means a Channel Buffer can hold up to 3 values
	msgQueue := make(chan string, 3)

	msgQueue <- "One"
	msgQueue <- "Two"
	msgQueue <- "Three"

	// Even though we close this non-empty Channel, we can still receive
	// the remaining values (see below)
	close(msgQueue)

	// We use the range keyword to iterate over each element as it gets
	// received from the msgQueue
	for m := range msgQueue {
		fmt.Println(m)
	}
}

package main

import "fmt"

func main() {
	// Created buffered channel of strings with a capacity of 3
	// This means a Channel Buffer can hold up to 3 values
	msgQueue := make(chan string, 3)

	msgQueue <- "One"
	msgQueue <- "Two"
	msgQueue <- "Three"

	// We drain the msgQueue by receiving all the values from the buffered channel
	fmt.Println(<-msgQueue)
	fmt.Println(<-msgQueue)
	fmt.Println(<-msgQueue)
}

package main

import "fmt"

func main() {
	c := make(chan string, 2) // 2 = Buffer
	c <- "Aditya"
	c <- "Hajare"
	// c <- "Test123" // Since channel is already full, this will cause deadlock

	fmt.Println(<-c)
	fmt.Println(<-c)
}

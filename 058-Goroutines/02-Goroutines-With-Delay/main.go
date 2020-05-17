package main

import (
	"fmt"
	"time"
)

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("नमस्ते आदित्य!", i, source)
	}
	time.Sleep(time.Millisecond * 1) // Sleep for 1 millisecond
}

func main() {
	go printGreetings("From Goroutine")
	printGreetings("From Main Function")
}

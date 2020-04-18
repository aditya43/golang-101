package main

import "fmt"

var done chan bool = make(chan bool)

func sayNamaste(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println("नमस्ते आदित्य!", i, source)
	}

	if source == "Goroutine" {
		done <- true
	}
}

func main() {
	go sayNamaste("Goroutine")
	sayNamaste("main() Function")

	<-done
}

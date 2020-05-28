package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1) // 1 because we will be executing 1 goroutine
	go func() {      // Goroutine
		printStrFiveTimes("नमस्ते आदित्य") // Call our func inside goroutine
		waitGroup.Done()                   // Notify waitGroup that our func is done executing
	}()

	waitGroup.Wait() // Wait for number of goroutines set under waitGroup
}

func printStrFiveTimes(str string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%2v %v\n", i, str)
		time.Sleep(time.Millisecond * 500)
	}
}

package main

import (
	"fmt"
	"time"
)

// Sending and receiving on a channel is a blocking operation.
func main() {
	channel := make(chan string)

	go printStrFiveTimes("नमस्ते आदित्य", channel)

	for output := range channel {
		fmt.Println(output)
	}
}

func printStrFiveTimes(str string, channel chan string) {
	for i := 1; i <= 5; i++ {
		txt := fmt.Sprintf("%2v %v\n", i, str)
		channel <- txt
		time.Sleep(time.Millisecond * 500)
	}

	// Always close a channel at sender's end
	close(channel) // If we dont close a channel, it will cause deadlock
}

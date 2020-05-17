// Execute this pragram using built-in race detector as follows:
// go run -race main.go

package main

import (
	"fmt"
	"sync"
)

var msg string
var done chan bool
var mutex = &sync.Mutex{}

func sayNamaste() {
	mutex.Lock()
	msg = "नमस्ते आदित्य! From sayNamaste() Function"
	mutex.Unlock()

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go sayNamaste()

	mutex.Lock()
	msg = "नमस्ते आदित्य! From main() Function"
	fmt.Println(msg)
	mutex.Unlock()

	<-done // Making a receive call on "done" channel
}

// Execute this pragram using built-in race detector as follows:
// go run -race main.go

package main

import "fmt"

var msg string
var done chan bool

func sayNamaste() {
	msg = "नमस्ते आदित्य! From sayNamaste() Function"
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go sayNamaste()

	msg = "नमस्ते आदित्य! From main() Function"
	fmt.Println(msg)

	<-done // Making a receive call on "done" channel
}

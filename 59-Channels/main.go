package main

import "fmt"

func sayNamasteFromChannel(msg chan string) {
	msg <- "नमस्ते आदित्य!"
}

func main() {
	msgFromChannel := make(chan string)
	fmt.Printf("msgFromChannel: %#v\n", msgFromChannel)

	go sayNamasteFromChannel(msgFromChannel)
	fmt.Println("Message from Channel:", <-msgFromChannel)

	close(msgFromChannel)
}

package main

import "fmt"

func sayNamasteFromChannel(msg chan string, msgType int) {
	if msgType == 1 {
		msg <- "नमस्ते आदित्य!"
	} else {
		msg <- "नमस्ते निशि!"
	}
}

func main() {
	msgFromChannel := make(chan string)
	fmt.Printf("msgFromChannel: %#v\n", msgFromChannel)

	go sayNamasteFromChannel(msgFromChannel, 1)
	fmt.Println("Message from Channel:", <-msgFromChannel)

	go sayNamasteFromChannel(msgFromChannel, 2)
	fmt.Println("Message from Channel:", <-msgFromChannel)

	fmt.Println("Inside main() function")
	close(msgFromChannel)
}

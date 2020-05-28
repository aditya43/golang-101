package main

import (
	"fmt"
	"time"
)

// Concurrency allows parallelism
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			go Every500ms(c1)
		}
	}()

	go func() {
		for {
			go Every2Seconds(c2)
		}
	}()

	// This will print 1 after another.. NEVER EVER DO THIS!
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// Correct approach
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func Every500ms(c1 chan string) {
	c1 <- "Every 500ms"
	time.Sleep(time.Millisecond * 500)
}

func Every2Seconds(c2 chan string) {
	c2 <- "Every 2 Seconds"
	time.Sleep(time.Second * 2)
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// This will automatically close connection after 10 seconds
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "Input: %s\n", line)
	}

	defer conn.Close()

	// Code will reach this line..
	// We have an open stream connection
	// How does the above reader knows when it's done?
	fmt.Println("Code will reach this line")
}

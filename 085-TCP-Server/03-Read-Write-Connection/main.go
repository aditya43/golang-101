package main

import (
	"bufio"
	"fmt"
	"net"
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
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "Input: %s\n", line)
	}

	defer conn.Close()

	// Code will never reach this line..
	// We have an open stream connection
	// How does the above reader knows when it's done?
	fmt.Println("Code will never reach this line")
}

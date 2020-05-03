package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}

		_, _ = io.WriteString(conn, "\nHello Aditya\n")
		fmt.Fprintln(conn, "How are you?")
		fmt.Fprintf(conn, "%v", "Nishi here..")

		conn.Close()
	}
}

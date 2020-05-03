package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
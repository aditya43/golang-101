package main

import (
	"time"

	"github.com/aditya43/golang-core/120-Functional-Options-Pattern/server"
)

func main() {
	srv := server.New(
		server.WithHost("localhost"),
		server.WithMaxConn(100),
		server.WithPort(8080),
		server.WithTimeout(time.Minute),
	)

	srv.Start()
	time.Sleep(time.Second)
	srv.Stop()
}

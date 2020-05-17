package main

import (
	"fmt"

	"github.com/aditya43/golang/20-Slices-In-Go/38-Limiting-The-Backing-Array-Sharing/api"
)

func main() {
	received := api.Read(0, 3)

	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}

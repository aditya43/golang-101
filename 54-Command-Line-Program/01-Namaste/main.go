package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var err error
	var repeatCount int

	args := os.Args[1:]

	if len(args) < 1 {
		err = errors.New("Please specify integer count")
	} else {
		repeatCount, err = strconv.Atoi(args[0])
	}

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < repeatCount; i++ {
		fmt.Println("नमस्ते आदित्य")
	}
}

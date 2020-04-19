package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://postman-echo.com/get?foo1=bar1&foo2=bar2")

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body)) // Without pretty print JSON

	_ = err
}



package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func fetchUrl(url string, wg *sync.WaitGroup) {

	// Decrement the WaitGroup counter once we've fetched the URL
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch the URL, ", url, " and encountered this error: ", err)
	} else {
		fmt.Println("Contents of url, ", url, ", is:")
		contents, err := ioutil.ReadAll(response.Body)

		response.Body.Close()
		if err != nil {
			log.Fatal("Failed to read the response body and encountered this error: ", err)
		}

		output := &bytes.Buffer{}
		if err := json.Indent(output, contents, "", "  "); err != nil {
			panic(err)
		}

		fmt.Println(output.String())
		// fmt.Println(string(contents)) // Without pretty printing JSON
	}

}

func main() {

	var wg sync.WaitGroup
	var urlList = []string{
		"https://reqres.in/api/users/1",
		"https://reqres.in/api/users/2",
		"https://reqres.in/api/users/3",
	}

	// Loop through the list of URLs
	for _, url := range urlList {
		// Increment the WaitGroup counter
		wg.Add(1)
		// Call the fetchURL function as a goroutine to fetch the URL
		go fetchUrl(url, &wg)
	}
	// Wait for the goroutines that are part of the WaitGroup to finish
	wg.Wait()
}

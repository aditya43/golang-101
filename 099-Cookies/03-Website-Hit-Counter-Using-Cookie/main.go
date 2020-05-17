package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(res http.ResponseWriter, req *http.Request) {
	hits := updateHits(res, getExistingHits(req))

	_, _ = io.WriteString(res, `Hits: `+hits)
}

// Get current hits from hits counter cookie otherwise return 0
func getExistingHits(req *http.Request) int {
	c, err := req.Cookie("adi-hits")

	if err != nil {
		// Cookie not found or there was an err
		fmt.Println(err.Error())
		return 0
	}

	hits, err := strconv.Atoi(c.Value)

	if err != nil {
		// Wrong value in cookie or error converting it to int
		fmt.Println(err.Error())
		return 0
	}

	return hits
}

// Set adi-hits cookie with updated hits
func updateHits(res http.ResponseWriter, hits int) string {
	val := strconv.Itoa(hits + 1)

	http.SetCookie(res, &http.Cookie{
		Name:  "adi-hits",
		Value: val,
	})

	return val
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(res, "In the terminal")
}

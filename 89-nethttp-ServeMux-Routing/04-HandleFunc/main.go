package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	_ = http.ListenAndServe(":8080", nil)
}

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	log.Fatalln(http.ListenAndServe(":80", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "नमस्ते आदित्य on AWS.")
}

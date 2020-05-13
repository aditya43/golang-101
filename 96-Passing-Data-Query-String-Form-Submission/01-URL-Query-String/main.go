package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", adi)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func adi(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("uname")
	_, _ = io.WriteString(res, "नमस्ते "+name)
}

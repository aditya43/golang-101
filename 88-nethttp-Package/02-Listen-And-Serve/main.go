package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code we want in this func")
}

func main() {
	var d hotdog
	_ = http.ListenAndServe(":8080", d)
}

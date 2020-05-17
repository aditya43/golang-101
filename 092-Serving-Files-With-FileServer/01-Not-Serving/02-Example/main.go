package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	_ = http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, _ = io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/toby.jpg">
	`)
}

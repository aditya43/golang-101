package main

import (
	"io"
	"net/http"
)

func main() {
	handler := http.StripPrefix("/resources", http.FileServer(http.Dir("./assets")))

	http.HandleFunc("/dog", dog)
	http.Handle("/resources/", handler)

	_ = http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8;")
	_, _ = io.WriteString(res, `<img src="/resources/toby.jpg">`)
}

package main

import (
	"io"
	"log"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		// _, _ = io.WriteString(res, "404 Not Found")
		http.NotFound(res, req)
		return
	}

	_, _ = io.WriteString(res, "Home Page")
}

func login(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, "Login Page")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)

	// Passing "nil" to "ListenAndServe" will use default "ServeMux"
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

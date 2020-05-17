package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/http-error", httpError)

	log.Fatal(http.ListenAndServe(":8080", nil)) // log.Fatal
}

func httpError(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Custom-Header", "Aditya Hajare")
	http.Error(res, "Custom not found error", http.StatusNotFound)
	// OR
	// http.Error(res, "Custom not found error", 404)
}

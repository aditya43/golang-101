package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aditya43/golang/65-Web-Server/validationkit"
)

func sayNamasteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "नमस्ते आदित्य!")
}

func checkUsernameSyntaxHandler(w http.ResponseWriter, r *http.Request) {
	var usernameSyntaxResult bool
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username not provided!", http.StatusInternalServerError)
	} else {
		usernameSyntaxResult = validationkit.CheckUsernameSyntax(username)
		fmt.Fprintf(w, "Syntax Check Result for %v is %v", username, usernameSyntaxResult)
	}
}

func main() {
	http.HandleFunc("/say-namaste", sayNamasteHandler)
	http.HandleFunc("/username-syntax-checker", checkUsernameSyntaxHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

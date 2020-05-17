package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/set-cookie", set)
	http.HandleFunc("/get-cookie", get)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", home)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "aditya",
		Value: "Aditya Testing",
		Path:  "/",
	})

	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK IN BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}

func get(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("aditya")

	if err != nil {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, "Cookie: ", cookie)
}

func home(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, `
		<a href="/get-cookie">Get Cookie</a><br>
		<a href="/set-cookie">Set Cookie</a>
	`)
}

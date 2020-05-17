package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", aditya)
	_ = http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(w, "bar ran")
}

func aditya(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "Aditya")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

// Type representing a gopher
type Gopher struct {
	Name string
}

// Handler for the namaste-user route
func namasteGopherHandler(w http.ResponseWriter, r *http.Request) {
	var userName string

	userName = r.URL.Query().Get("userName")

	if userName == "" {
		userName = "आदित्य" // Default username
	}

	gopher := Gopher{Name: userName}

	renderTemplate(w, "./templates/namaste.html", gopher)
}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)

	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}

	_ = t.Execute(w, templateData)
}

func main() {
	http.HandleFunc("/namaste-user", namasteGopherHandler)

	_ = http.ListenAndServe(":8080", nil)
}

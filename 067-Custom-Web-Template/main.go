// Creating and using a template in Go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/aditya43/golang/57-Social-Media-Post-Custom-Type/socialmedia"
)

// Handler for displaying a social media post
func displaySocialMediaPostHandler(w http.ResponseWriter, r *http.Request) {
	myPost := socialmedia.NewPost(
		"Aditya Hajare",
		socialmedia.Moods["thrilled"],
		"नमस्ते आदित्य",
		"नमस्ते आदित्य! A sample post",
		"http://adiinviter.com",
		"",
		"",
		[]string{"नमस्ते", "आदित्य", "नमस्ते आदित्य a sample post"},
	)

	fmt.Printf("myPost: %+v\n", myPost)
	renderTemplate(w, "./templates/socialmediapost.html", myPost)
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
	http.HandleFunc("/display-social-media-post", displaySocialMediaPostHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server is listening on http://localhost:8080")

	_ = http.ListenAndServe(":8080", nil)
}

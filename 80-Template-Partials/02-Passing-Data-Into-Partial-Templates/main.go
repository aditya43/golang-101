package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.adi"))
	tpl, _ = tpl.ParseGlob("templates/partials/*.adi")
}

func main() {
	adi := struct {
		Name string
		Age  int
	}{
		"Aditya Hajare",
		32,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "main.adi", adi)

	if err != nil {
		log.Fatalln(err)
	}
}

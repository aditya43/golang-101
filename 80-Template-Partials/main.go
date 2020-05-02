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
	err := tpl.ExecuteTemplate(os.Stdout, "main.adi", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

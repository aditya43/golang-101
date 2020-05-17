package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

// Always make sure that our templates are parsed only once
func init() {
	//Parsing all templates in "/templates" directory
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "one.adi", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("\n\n-----------------------------\n\n")

	err = tpl.ExecuteTemplate(os.Stdout, "two.adi", nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("\n\n-----------------------------\n\n")

	err = tpl.ExecuteTemplate(os.Stdout, "three.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

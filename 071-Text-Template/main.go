package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("adi.anything", "adi.txt")

	if err != nil {
		log.Fatalln(err)
		return
	}

	// This will execute 1st parsed template from the container
	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n\n---------------------------\n\n")

	// To execute a specific template from parsed templates container
	err = tpl.ExecuteTemplate(os.Stdout, "adi.txt", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
)

func passStruct() {
	type Person struct {
		Name string
		Age  int
	}

	adi := Person{
		"Aditya",
		32,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "03-struct.adi", adi)

	if err != nil {
		log.Fatalln(err)
	}
}

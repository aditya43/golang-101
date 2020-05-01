package main

import (
	"log"
	"os"
)

func passSliceOfStruct() {
	type Person struct {
		Name string
		Age  int
	}

	adi := Person{"Aditya", 32}
	nishi := Person{"Nishi", 26}

	users := []Person{adi, nishi}

	err := tpl.ExecuteTemplate(os.Stdout, "04-slice-struct.adi", users)

	if err != nil {
		log.Fatalln(err)
	}
}
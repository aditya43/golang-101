package main

import (
	"log"
	"os"
)

func passSlice() {
	names := []string{"Aditya", "Nishi", "John", "Jane"}

	err := tpl.ExecuteTemplate(os.Stdout, "01-slice.adi", names)

	if err != nil {
		log.Fatalln(err)
	}
}

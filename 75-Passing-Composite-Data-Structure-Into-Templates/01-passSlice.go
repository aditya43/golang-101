package main

import (
	"fmt"
	"log"
	"os"
)

func passSlice() {
	fmt.Print("\n-------------------------------------------\n")
	fmt.Print("Slice")
	fmt.Print("\n-------------------------------------------\n")

	names := []string{"Aditya", "Nishi", "John", "Jane"}

	err := tpl.ExecuteTemplate(os.Stdout, "01-slice.adi", names)

	if err != nil {
		log.Fatalln(err)
	}
}

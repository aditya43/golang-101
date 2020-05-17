package main

import (
	"log"
	"os"
)

func passMap() {
	users := map[int]string{
		111: "Aditya",
		222: "Nishi",
		333: "John",
		444: "Jane",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "02-map.adi", users)

	if err != nil {
		log.Fatalln(err)
	}
}

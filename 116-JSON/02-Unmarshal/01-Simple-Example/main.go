package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Website struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"website"`
}

func main() {
	var p person

	jsonStr := `"{"name":"Aditya","city":"Pune","website":{"name":"Github","url":"https://github.com/aditya43"}}"`

	err := json.Unmarshal([]byte(jsonStr), &p)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(p.Name)
	fmt.Println(p.City)
	fmt.Println(p.Website.Name)
	fmt.Println(p.Website.Url)
}

package main

import (
	"log"
	"os"
)

func passStructOfSliceOfStruct() {
	type Person struct {
		name string
		age  int
	}

	type Car struct {
		model string
		price int
	}

	type Data struct {
		People []Person
		Cars   []Car
	}

	adi := Person{"Aditya", 32}
	nishi := Person{"Nishi", 26}

	merc := Car{"Mercedes", 4000000}
	por := Car{"Porche", 9000000}

	people := []Person{adi, nishi}
	cars := []Car{merc, por}

	data := Data{people, cars}

	err := tpl.ExecuteTemplate(os.Stdout, "05-struct-slice-struct.adi", data)

	if err != nil {
		log.Fatalln(err)
	}
}

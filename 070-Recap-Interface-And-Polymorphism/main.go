package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Developer struct {
	Person
	lang string
}

func (p Person) details() {
	fmt.Printf("%s is %d years old.\n", p.name, p.age)
}

func (d Developer) details() {
	fmt.Printf("%s is %d years old and can code in %s language(s).\n", d.name, d.age, d.lang)
}

// Interface - Person and Developer implicitely implements this interface since they both have details() function.
type Human interface {
	details()
}

// This function can call details() method from any type implementing Human interface.
func details(h Human) {
	h.details()
}

func main() {
	adi := Person{
		"Aditya Hajare",
		32,
	}

	dev := Developer{
		adi,
		"Go Lang, JavaScript, Python, C++, PHP, Java, C",
	}

	adi.details()
	dev.details()

	details(adi)
	details(dev)
}

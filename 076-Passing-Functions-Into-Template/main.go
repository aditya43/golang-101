package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// Create a "FuncMap" to register functions.
// "uc" is what the func will be called in template. Its a "ToUpper" func from package "strings".
// "ft" is a custom func and it slices a string and returns first 3 characters.
var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThreeChars,
}

// Return first 3 characters
func firstThreeChars(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/adi.gohtml"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "adi.gohtml", data)

	if err != nil {
		log.Fatalln(err)
	}
}

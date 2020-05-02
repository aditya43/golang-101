package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

var funcMap = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("adi.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "adi.gohtml", time.Now())

	if err != nil {
		log.Fatalln(err)
	}
}

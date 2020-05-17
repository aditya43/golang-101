package main

import "text/template"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	passSlice()
	passMap()
	passStruct()
	passSliceOfStruct()
	passStructOfSliceOfStruct()
}

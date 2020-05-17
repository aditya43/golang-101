package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, WebAppRoot+"/templates/header.html", WebAppRoot+"/templates/footer.html")
	if err != nil {
		log.Println("Error encountered while parsing the template: ", err)
	}
	_ = t.Execute(w, templateData)
}

func RenderGatedTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, WebAppRoot+"/templates/gatedheader.html", WebAppRoot+"/templates/footer.html")
	if err != nil {
		log.Println("Error encountered while parsing the template: ", err)
	}
	_ = t.Execute(w, templateData)
}

/*
func RenderUnsafeTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := ttemplate.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	w.Header().Set("X-XSS-Protection", "0")
	t.Execute(w, templateData)
}
*/

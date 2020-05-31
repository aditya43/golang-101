package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	c = appendData(w, c)
	vals := strings.Split(c.Value, "|")
	_ = tpl.ExecuteTemplate(w, "index.gohtml", vals)
}

// add func to get cookie
func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

// Store multiple pic filenames in cookie
func appendData(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	p1 := "aditya.jpg"
	p2 := "nishi.jpg"

	s := c.Value

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}

	if !strings.Contains(s, p2) {
		s += "|" + p2
	}

	c.Value = s
	http.SetCookie(w, c)
	return c
}

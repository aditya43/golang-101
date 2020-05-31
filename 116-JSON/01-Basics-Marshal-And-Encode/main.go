package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname  string
	Lname  string
	Cities []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := person{
		Fname:  "Aditya",
		Lname:  "Hajare",
		Cities: []string{"Pune", "Mumbai", "Delhi"},
	}

	err := json.NewEncoder(w).Encode(p)

	if err != nil {
		log.Fatalln(err)
	}
}

func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := person{
		Fname:  "Aditya",
		Lname:  "Hajare",
		Cities: []string{"Pune", "Mumbai", "Delhi"},
	}

	data, err := json.Marshal(p)

	if err != nil {
		log.Fatalln(err)
	}

	_, _ = w.Write(data)
}

func index(w http.ResponseWriter, r *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Home</title>
		</head>
		<body>
			Home Page
		</body>
		</html>`
	_, _ = w.Write([]byte(s))
}

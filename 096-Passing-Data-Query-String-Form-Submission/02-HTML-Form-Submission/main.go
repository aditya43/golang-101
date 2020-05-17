package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", adi)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func adi(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("uname")

	res.Header().Set("Content-Type", "text/html; charset=utf-8;")

	_, _ = io.WriteString(res, `
	<form method="POST">
		<input type="text" name="uname">
		<button type="submit">Submit</button>
	</form>
	<br /><hr><br />
	Input Value: ` + name)
}
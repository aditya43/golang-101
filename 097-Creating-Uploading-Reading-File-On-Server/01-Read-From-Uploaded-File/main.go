package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", adi)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func adi(resp http.ResponseWriter, req *http.Request) {
	var str string

	if req.Method == http.MethodPost {
		// Open
		file, headers, err := req.FormFile("myfile")

		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		// For our information
		fmt.Printf("File: %v\nHeaders: %v,Errors: %v\n", file, headers, err)

		// Read
		byteStream, err := ioutil.ReadAll(file)

		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		str = string(byteStream)
	}

	resp.Header().Set("Content-Type", "text/html; charset=utf-8;")

	cnt, err := io.WriteString(resp, `
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="myfile">
		<button type="submit">Submit</button>
		</form>
		<br><hr><br>
		File Contents: `+str)

	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Count: ", cnt)
}

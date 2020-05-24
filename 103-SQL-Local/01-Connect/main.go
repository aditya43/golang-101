package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:adi123@tcp(127.0.0.1:3306)/adi_test_db?charset=utf8")
	check(err)

	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	check(http.ListenAndServe(":8080", nil))
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "Successfully connected mysql database!")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

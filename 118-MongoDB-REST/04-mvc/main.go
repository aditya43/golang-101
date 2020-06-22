package main

import (
	"log"
	"net/http"

	"github.com/aditya43/golang-core/118-mongodb-rest/04-mvc/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}

package main

import (
	"log"
	"net/http"

	"github.com/aditya43/golang-core/118-MongoDB-REST/06-Storing-Data-In-Map-Instead-of-MongoDB/controllers"
	"github.com/aditya43/golang-core/118-MongoDB-REST/06-Storing-Data-In-Map-Instead-of-MongoDB/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}

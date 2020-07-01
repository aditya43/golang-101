package main

import (
	"log"
	"net/http"

	"github.com/aditya43/golang-core/118-MongoDB-REST/07-Storing-Data-In-JSON-File-Instead-of-MongoDB/controllers"
	"github.com/aditya43/golang-core/118-MongoDB-REST/07-Storing-Data-In-JSON-File-Instead-of-MongoDB/models"
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
	return models.LoadUsers()
}

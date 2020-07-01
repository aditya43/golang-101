package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aditya43/golang-core/118-MongoDB-REST/07-Storing-Data-In-JSON-File-Instead-of-MongoDB/models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	_ = json.NewDecoder(r.Body).Decode(&u)

	// create ID
	u.Id = uuid.NewV4().String()

	// store the user
	uc.session[u.Id] = u
	models.StoreUsers(uc.session)

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	models.StoreUsers(uc.session)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user ", id, "\n")
}

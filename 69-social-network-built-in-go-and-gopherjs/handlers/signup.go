package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/validationkit"
)

type SignUpForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplaySignUpForm displays the Sign Up form
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, WebAppRoot+"/templates/signupform.html", s)
}

func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, WebAppRoot+"/templates/signupconfirmation.html", s)
}

func PopulateFormFields(r *http.Request, s *SignUpForm) {

	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, e *common.Env) {

	PopulateFormFields(r, s)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "The username field is required."
	}

	// Check if first name was filled out
	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "The first name field is required."
	}

	// Check if last name was filled out
	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "The last name field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "The e-mail address field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "The password field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	// Check username syntax
	if !validationkit.CheckUsernameSyntax(r.FormValue("username")) {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	// Check e-mail address syntax
	if !validationkit.CheckEmailSyntax(r.FormValue("email")) {
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok {
			s.Errors["emailError"] += " " + emailErrorMessage
		} else {
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	// Check if password and confirm password field values match
	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "The password and confirm pasword fields do not match."
	}

	if len(s.Errors) > 0 {
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s, e)
	}

}

// ProcessSignUpForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, e *common.Env) {

	u := models.NewUser(r.FormValue("username"), r.FormValue("firstName"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("password"))
	//fmt.Println("user: ", u)
	err := e.DB.CreateUser(u)

	if err != nil {
		log.Print(err)
	}

	user, err := e.DB.GetUser("EngineerKamesh")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Fetch User Result: %+v\n", user)
	}

	// Display form confirmation message
	DisplayConfirmation(w, r, s)

}

func SignUpHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := SignUpForm{}
		s.FieldNames = []string{"username", "firstName", "lastName", "email"}
		s.Fields = make(map[string]string)
		s.Errors = make(map[string]string)
		s.PageTitle = "Sign Up"

		switch r.Method {

		case "GET":
			DisplaySignUpForm(w, r, &s)
		case "POST":
			ValidateSignUpForm(w, r, &s, e)
		default:
			DisplaySignUpForm(w, r, &s)
		}

	})
}

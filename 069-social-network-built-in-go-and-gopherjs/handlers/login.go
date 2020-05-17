package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/authenticate"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/utility"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/validationkit"
)

type LoginForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplayLoginForm displays the Sign Up form
func DisplayLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm) {
	fmt.Println("reached display login form")
	RenderTemplate(w, WebAppRoot+"/templates/loginform.html", l)
}

func PopulateLoginFormFields(r *http.Request, l *LoginForm) {

	for _, fieldName := range l.FieldNames {
		l.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateLoginForm validates the Sign Up form's fields
func ValidateLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {

	PopulateLoginFormFields(r, l)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		l.Errors["usernameError"] = "The username field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		l.Errors["passwordError"] = "The password field is required."
	}

	// Check username syntax
	if !validationkit.CheckUsernameSyntax(r.FormValue("username")) {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := l.Errors["usernameError"]; ok {
			l.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			l.Errors["usernameError"] = usernameErrorMessage
		}
	}

	if len(l.Errors) > 0 {
		DisplayLoginForm(w, r, l)
	} else {
		ProcessLoginForm(w, r, l, e)
	}

}

// ProcessLoginForm
func ProcessLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {

	authResult := authenticate.VerifyCredentials(e, r.FormValue("username"), r.FormValue("password"))
	fmt.Println("auth result: ", authResult)

	// Successful login, let's create a cookie for the user and redirect them to the feed route
	if authResult {

		sessionID := utility.GenerateUUID()
		fmt.Println("sessid: ", sessionID)
		u, err := e.DB.GetUser(r.FormValue("username"))
		if err != nil {
			log.Print("Encountered error when attempting to fetch user record: ", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		err = authenticate.CreateSecureCookie(u, sessionID, w, r)
		if err != nil {
			log.Print("Encountered error when attempting to create secure cookie: ", err)
			http.Redirect(w, r, "/login", 302)
			return

		}

		err = authenticate.CreateUserSession(u, sessionID, w, r)
		if err != nil {
			log.Print("Encountered error when attempting to create user session: ", err)
			http.Redirect(w, r, "/login", 302)
			return

		}

		http.Redirect(w, r, "/feed", 302)

	} else {

		l.Errors["usernameError"] = "Invalid login."
		DisplayLoginForm(w, r, l)

	}

}

func LoginHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		l := LoginForm{}
		l.FieldNames = []string{"username"}
		l.Fields = make(map[string]string)
		l.Errors = make(map[string]string)
		l.PageTitle = "Log In"

		switch r.Method {

		case "GET":
			DisplayLoginForm(w, r, &l)
		case "POST":
			ValidateLoginForm(w, r, &l, e)
		default:
			DisplayLoginForm(w, r, &l)
		}

	})
}

package handlers

import (
	"log"
	"net/http"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common/authenticate"

	"go.isomorphicgo.org/go/isokit"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common"
	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/forms"
)

func MyProfileHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		formParams := isokit.FormParams{ResponseWriter: w, Request: r}
		myProfileForm := forms.NewMyProfileForm(&formParams)
		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
		}

		uuid := gfSession.Values["uuid"]
		u, err := env.DB.GetUserProfile(uuid.(string))
		if err != nil {
			log.Print(err)
		}
		u.Username = gfSession.Values["username"].(string)
		u.Form = myProfileForm
		u.PageTitle = "My Profile"

		env.TemplateSet.Render("myprofile_page", &isokit.RenderParams{Writer: w, Data: u})
	})
}

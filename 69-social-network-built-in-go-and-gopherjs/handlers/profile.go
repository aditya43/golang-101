package handlers

import (
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
	"github.com/gorilla/mux"
	"go.isomorphicgo.org/go/isokit"
)

func ProfileHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		username := vars["username"]

		u, err := env.DB.GetGopherProfile(username)
		if err != nil {
			log.Print(err)
		}

		u.PageTitle = u.Username
		env.TemplateSet.Render("gopherprofile_page", &isokit.RenderParams{Writer: w, Data: u})

	})
}

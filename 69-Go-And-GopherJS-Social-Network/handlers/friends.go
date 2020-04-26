package handlers

import (
	"net/http"

	"go.isomorphicgo.org/go/isokit"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common"
)

func FriendsHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render("friends_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}

package handlers

import (
	"net/http"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}

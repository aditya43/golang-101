package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/authenticate"
)

func GatedContentHandler(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		shouldRedirectToLogin := false

		secureCookieMap, err := authenticate.ReadSecureCookieValues(w, r)

		log.Printf("secure cookie map: %+v", secureCookieMap)

		if err != nil {
			log.Print(err)
		}

		//fmt.Printf("secure cookie contents: %+v\n", secureCookieMap)

		// Check if the sid key which is used to store the session id value
		// has been populated in the map using the comma ok idiom
		if _, ok := secureCookieMap["sid"]; ok == true {

			gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")

			fmt.Printf("gopherface session values: %+v\n", gfSession.Values)
			if err != nil {
				log.Print(err)
				return
			}

			// Check if the session id stored in the secure cookie matches
			// the id and username on the server-side session
			if gfSession.Values["sessionID"] == secureCookieMap["sid"] && gfSession.Values["username"] == secureCookieMap["username"] {
				next(w, r)
			} else {
				shouldRedirectToLogin = true
			}

		} else {
			shouldRedirectToLogin = true

		}

		if shouldRedirectToLogin == true {
			http.Redirect(w, r, "/login", 302)
		}

	})

}

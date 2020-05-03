package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/authenticate"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"
)

func SaveUserProfileEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"]

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		var u models.UserProfile
		json.Unmarshal(reqBody, &u)

		env.DB.UpdateUserProfile(uuid.(string), u.About, u.Location, u.Interests)

	})
}

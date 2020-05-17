package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/authenticate"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
)

func FindGophersEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		var searchTerm string
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		err = json.Unmarshal(reqBody, &searchTerm)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		gophers, err := env.DB.FindGophers(uuid, searchTerm)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gophers)

	})
}

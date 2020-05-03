package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
)

func GetGopherProfileEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var gopherUUID string
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		err = json.Unmarshal(reqBody, &gopherUUID)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		u, err := env.DB.GetGopherProfile(gopherUUID)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)

	})
}

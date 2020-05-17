package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models/socialmedia"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/authenticate"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
)

func SavePostEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		var p socialmedia.Post
		json.Unmarshal(reqBody, &p)

		err = env.DB.SavePost(uuid, p.Caption, p.MessageBody, p.RawMoodValue)

		if err != nil {
			w.Write([]byte("ERROR: Failed to save post!"))
		} else {
			w.Write([]byte("Post saved successfully!"))
		}

	})
}

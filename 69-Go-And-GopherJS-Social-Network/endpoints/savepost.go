package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/models/socialmedia"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common/authenticate"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common"
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

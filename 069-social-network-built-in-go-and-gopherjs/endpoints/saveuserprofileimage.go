package endpoints

import (
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/authenticate"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common/utility"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/tasks"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/common"
)

func SaveUserProfileImageEndpoint(env *common.Env) http.HandlerFunc {
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

		contentType := http.DetectContentType(reqBody)
		extensions, err := mime.ExtensionsByType(contentType)

		if err != nil {
			log.Print(err)
		}

		var extension string

		if len(extensions) > 0 {
			extension = extensions[0]
		}

		if extension != "" {

			// Remove the existing profile image and its thumbnail (if they exist)
			u, err := env.DB.GetUserProfile(uuid)

			if err != nil {
				log.Println(err)
			}

			if u.ProfileImagePath != "" {
				os.Remove(WebAppRoot + u.ProfileImagePath)
				os.Remove(strings.Replace(WebAppRoot+u.ProfileImagePath, "_thumb", "", -1))

			}

			randomFileName := utility.GenerateUUID()
			imageFilePathWithoutExtension := "/static/uploads/images/" + randomFileName
			err = ioutil.WriteFile(WebAppRoot+imageFilePathWithoutExtension+extension, reqBody, 0666)
			if err != nil {
				log.Println(err)
			}

			thumbnailResizeTask := tasks.NewImageResizeTask(WebAppRoot+imageFilePathWithoutExtension, extension)
			thumbnailResizeTask.Perform()

			// Update the database record to the path of the image file
			env.DB.UpdateUserProfileImage(uuid, imageFilePathWithoutExtension+"_thumb"+extension)

			w.Write([]byte(imageFilePathWithoutExtension + "_thumb" + extension))

		} else {
			w.Write([]byte("Error: Failed to process uploaded file."))

		}

	})
}

package handlers

import (
	"context"
	"encoding/json"
	"strings"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/client/common"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/forms"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"
	"github.com/gopherjs/gopherjs/js"

	"go.isomorphicgo.org/go/isokit"
)

var myProfileForm *forms.MyProfileForm

func MyProfileHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		myProfileForm = forms.NewMyProfileForm(nil)
		var profile chan models.UserProfile = make(chan models.UserProfile, 1)
		GetUserProfileRequest(profile)
		u := <-profile
		u.Form = myProfileForm
		u.PageTitle = "My Profile"
		env.TemplateSet.Render("myprofile_content", &isokit.RenderParams{Data: u, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})
		InitializeMyProfileEventHandlers(env)
	})
}

func InitializeMyProfileEventHandlers(env *common.Env) {

	if myProfileForm == nil {
		myProfileForm = forms.NewMyProfileForm(nil)
	}

	saveProfileButton := env.Document.GetElementByID("saveProfileButton").(*dom.HTMLButtonElement)
	saveProfileButton.AddEventListener("click", false, func(event dom.Event) {
		VerifyProfileForm(env, event, myProfileForm)
	})

	uploadInput := env.Document.GetElementByID("imagefile").(*dom.HTMLInputElement)
	uploadInput.AddEventListener("change", false, func(event dom.Event) {
		go UploadProfileImageRequest(env, event)
	})

	go PopulateFriendsList(env)
}

func blobToBytes(blob *js.Object) []byte {
	var b = make(chan []byte)
	fileReader := js.Global.Get("FileReader").New()
	fileReader.Set("onload", func() {
		b <- js.Global.Get("Uint8Array").New(fileReader.Get("result")).Interface().([]byte)
	})
	fileReader.Call("readAsArrayBuffer", blob)
	return <-b
}

func UploadProfileImageRequest(env *common.Env, event dom.Event) {
	fileElement := event.Target()
	file := fileElement.Underlying().Get("files").Index(0)
	b := blobToBytes(file)

	data, err := xhr.Send("POST", "/restapi/save-user-profile-image", b)
	if err != nil {
		println(err)
	}

	if !strings.Contains(string(data), "Error:") {
		profileImage := env.Document.GetElementByID("profileImage").(*dom.HTMLImageElement)
		profileImage.SetAttribute("src", string(data))
		js.Global.Get("alertify").Call("success", "Profile image changed successfully!")

	} else {
		js.Global.Get("alertify").Call("error", string(data))
	}
}

func VerifyProfileForm(env *common.Env, event dom.Event, profileForm *forms.MyProfileForm) {

	event.PreventDefault()

	formElement := env.Document.GetElementByID("userProfileForm").(*dom.HTMLFormElement)
	profileForm.SetFormParams(&isokit.FormParams{FormElement: formElement})
	validationResult := profileForm.Validate()
	if validationResult == true {

		about := env.Document.GetElementByID("aboutTextArea").(*dom.HTMLTextAreaElement).Value
		location := env.Document.GetElementByID("locationInput").(*dom.HTMLInputElement).Value
		interests := env.Document.GetElementByID("interestsInput").(*dom.HTMLInputElement).Value

		u := models.UserProfile{}
		u.About = about
		u.Location = location
		u.Interests = interests
		go SaveProfileRequest(u)

	} else {
		profileForm.DisplayErrors()
	}

}

func SaveProfileRequest(u models.UserProfile) {

	profileBytes, err := json.Marshal(u)
	if err != nil {
		println(err)
	}
	data, err := xhr.Send("POST", "/restapi/save-user-profile", profileBytes)
	if err != nil {
		println(err)
	}

	if strings.Contains(string(data), "Error:") {

		js.Global.Get("alertify").Call("error", string(data))

	} else {

		js.Global.Get("alertify").Call("success", "Profile saved successfully!")
	}

}

func GetUserProfileRequest(profile chan models.UserProfile) {

	data, err := xhr.Send("GET", "/restapi/get-user-profile", nil)
	if err != nil {
		println("Encountered error while making XHR call: ", err)
	}

	var u models.UserProfile
	err = json.Unmarshal(data, &u)
	if err != nil {
		println("Encountered error while unmarshalling JSON: ", err)
	}

	profile <- u

}

func PopulateFriendsList(env *common.Env) {

	data, err := xhr.Send("POST", "/restapi/get-friends-list", nil)
	if err != nil {
		println("Encountered error while attempting to submit POST request via XHR: ", err)
		println(err)
	}

	var gophers []models.Gopher
	json.Unmarshal(data, &gophers)

	friendsListContainer := env.Document.GetElementByID("friendsListContainer").(*dom.HTMLDivElement)
	env.TemplateSet.Render("partials/friends_list", &isokit.RenderParams{Data: gophers, Disposition: isokit.PlacementReplaceInnerContents, Element: friendsListContainer})

	unfollowButtons := friendsListContainer.QuerySelectorAll(".unfollowButton")
	for i := 0; i < len(unfollowButtons); i++ {
		unfollowButtons[i].AddEventListener("click", false, func(event dom.Event) {
			go UnfollowGopherRequest(env, event)
		})
	}

}

func UnfollowGopherRequest(env *common.Env, event dom.Event) {
	button := event.Target()
	button.SetAttribute("disabled", "")
	uuid := button.GetAttribute("data-uuid")
	username := button.GetAttribute("data-username")
	friendItemDiv := env.Document.GetElementByID("followingitem_" + uuid).(*dom.HTMLDivElement)
	textBytes, err := json.Marshal(uuid)

	if err != nil {
		println("Encountered error while attempting to marshal JSON: ", err)
		println(err)
	}

	data, err := xhr.Send("POST", "/restapi/unfollow-gopher", textBytes)
	if err != nil {
		println("Encountered error while attempting to submit POST request via XHR: ", err)
		println(err)
	}

	var resultErr error
	json.Unmarshal(data, &resultErr)

	if resultErr == nil {
		js.Global.Get("alertify").Call("success", "Unfollowed "+username+" successfully!")
		friendItemDiv.ParentNode().RemoveChild(friendItemDiv)

	} else {
		js.Global.Get("alertify").Call("error", "Unfollow operation failed.")
		button.RemoveAttribute("disabled")
	}

}

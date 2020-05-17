package handlers

import (
	"context"
	"encoding/json"
	"strconv"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/client/common"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"
	"github.com/gopherjs/gopherjs/js"

	"go.isomorphicgo.org/go/isokit"
)

const ENTERKEY int = 13

func FriendsHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		println("Client-side Friends Handler")
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render("friends_content", &isokit.RenderParams{Data: m, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})
		InitializeFriendsEventHandlers(env)
	})

}

func InitializeFriendsEventHandlers(env *common.Env) {

	friendSearchButton := env.Document.GetElementByID("friendSearchButton").(*dom.HTMLButtonElement)
	friendSearchButton.AddEventListener("click", false, func(event dom.Event) {
		go FriendSearchRequest(env, event, true)
	})

	friendSearchInput := env.Document.GetElementByID("friendSearchInput").(*dom.HTMLInputElement)
	friendSearchInput.AddEventListener("keypress", false, func(event dom.Event) {
		if event.Underlying().Get("keyCode").Int() == ENTERKEY {
			event.PreventDefault()
			go FriendSearchRequest(env, event, true)
		}
	})

}

func FriendSearchRequest(env *common.Env, event dom.Event, showAlerts bool) {

	friendSearchInput := env.Document.GetElementByID("friendSearchInput").(*dom.HTMLInputElement)
	searchTerm := friendSearchInput.Value

	if searchTerm == "" {
		return
	}

	textBytes, err := json.Marshal(searchTerm)
	if err != nil {
		println("Encountered error while attempting to marshal JSON: ", err)
		println(err)
	}

	data, err := xhr.Send("POST", "/restapi/find-gophers", textBytes)
	if err != nil {
		println("Encountered error while attempting to submit POST request via XHR: ", err)
		println(err)
	}

	var gophers []models.Gopher
	json.Unmarshal(data, &gophers)

	if showAlerts == true {
		if len(gophers) > 0 {
			js.Global.Get("alertify").Call("success", "Found "+strconv.Itoa(len(gophers))+" gopher(s)!")
		} else {
			js.Global.Get("alertify").Call("error", "No gophers found.")
		}
	}

	searchResultsContainer := env.Document.GetElementByID("searchResultsContainer").(*dom.HTMLDivElement)
	env.TemplateSet.Render("partials/friend_search_results", &isokit.RenderParams{Data: gophers, Disposition: isokit.PlacementReplaceInnerContents, Element: searchResultsContainer})

	followButtons := searchResultsContainer.QuerySelectorAll(".followButton")
	for i := 0; i < len(followButtons); i++ {
		followButtons[i].AddEventListener("click", false, func(event dom.Event) {
			go FollowGopherRequest(env, event)
		})
	}

}

func FollowGopherRequest(env *common.Env, event dom.Event) {
	button := event.Target()
	button.SetAttribute("disabled", "")
	uuid := button.GetAttribute("data-uuid")
	textBytes, err := json.Marshal(uuid)

	if err != nil {
		println("Encountered error while attempting to marshal JSON: ", err)
		println(err)
	}

	data, err := xhr.Send("POST", "/restapi/follow-gopher", textBytes)
	if err != nil {
		println("Encountered error while attempting to submit POST request via XHR: ", err)
		println(err)
	}

	var resultErr error
	json.Unmarshal(data, &resultErr)

	if resultErr == nil {
		js.Global.Get("alertify").Call("success", "Followed successfully!")

	} else {
		js.Global.Get("alertify").Call("error", "Follow operation failed.")
	}

}

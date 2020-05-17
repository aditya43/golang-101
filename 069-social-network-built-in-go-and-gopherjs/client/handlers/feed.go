package handlers

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models/socialmedia"

	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/client/common"
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/forms"
	"github.com/gopherjs/gopherjs/js"

	"go.isomorphicgo.org/go/isokit"
)

var postForm *forms.SocialMediaPostForm

func FeedHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		println("Client-side Feed Handler")
		postForm = forms.NewSocialMediaPostForm(nil)
		postForm.PageTitle = "Feed"
		env.TemplateSet.Render("feed_content", &isokit.RenderParams{Data: postForm, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})

		InitializeFeedEventHandlers(env)
	})
}

func InitializeFeedEventHandlers(env *common.Env) {

	if postForm == nil {
		postForm = forms.NewSocialMediaPostForm(nil)
	}

	postButton := env.Document.GetElementByID("postButton").(*dom.HTMLButtonElement)
	postButton.AddEventListener("click", false, func(event dom.Event) {
		VerifyPostForm(env, event, postForm)
	})

	go PopulateFeedPosts(env)

}

func VerifyPostForm(env *common.Env, event dom.Event, postForm *forms.SocialMediaPostForm) {
	event.PreventDefault()

	formElement := env.Document.GetElementByID("smpostform").(*dom.HTMLFormElement)
	postForm.SetFormParams(&isokit.FormParams{FormElement: formElement})
	validationResult := postForm.Validate()
	if validationResult == true {

		caption := env.Document.GetElementByID("caption").(*dom.HTMLInputElement).Value
		messageBody := env.Document.GetElementByID("messageBody").(*dom.HTMLTextAreaElement).Value
		mood := env.Document.GetElementByID("mood").(*dom.HTMLSelectElement).Value

		p := socialmedia.Post{}
		p.Caption = caption
		p.MessageBody = messageBody
		moodValue, err := strconv.Atoi(mood)
		if err != nil {
			println(err)
			moodValue = 0
		}
		p.RawMoodValue = moodValue
		go SavePostRequest(env, p)
		formElement.Underlying().Call("reset")
	} else {
		postForm.DisplayErrors()
	}

}

func SavePostRequest(env *common.Env, p socialmedia.Post) {

	postBytes, err := json.Marshal(p)
	if err != nil {
		println(err)
	}
	data, err := xhr.Send("POST", "/restapi/save-post", postBytes)
	if err != nil {
		println(err)
	}

	if strings.Contains(string(data), "ERROR:") {

		js.Global.Get("alertify").Call("error", string(data))

	} else {

		js.Global.Get("alertify").Call("success", "Post saved successfully!")
		PopulateFeedPosts(env)
	}

}

func PopulateFeedPosts(env *common.Env) {

	data, err := xhr.Send("POST", "/restapi/fetch-posts", nil)
	if err != nil {
		println("Encountered error while attempting to submit POST request via XHR: ", err)
		println(err)
	}

	var posts []socialmedia.Post
	json.Unmarshal(data, &posts)

	for i := 0; i < len(posts); i++ {
		moodValue := posts[i].RawMoodValue
		posts[i].AuthorMoodEmoji = socialmedia.MoodsEmoji[socialmedia.MoodState(moodValue)]
		posts[i].TimeCreated = time.Unix(int64(posts[i].TimeCreatedUnixTS), 0)
	}

	postsContainer := env.Document.GetElementByID("postsContainer").(*dom.HTMLDivElement)
	env.TemplateSet.Render("partials/feed_posts", &isokit.RenderParams{Data: posts, Disposition: isokit.PlacementReplaceInnerContents, Element: postsContainer})

}

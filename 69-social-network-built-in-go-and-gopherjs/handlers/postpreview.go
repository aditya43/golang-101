package handlers

import (
	"net/http"
	"strconv"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models/socialmedia"
)

type PostForm struct {
	FieldNames  []string
	Fields      map[string]string
	Errors      map[string]string
	Moods       map[string]int
	MoodEmoji   map[string]string
	CurrentMood int
}

var MoodStates map[string]int
var MoodEmoji map[string]string

// DisplaySignUpForm displays the Sign Up form
func DisplayPostForm(w http.ResponseWriter, r *http.Request, p *PostForm) {
	RenderTemplate(w, WebAppRoot+"/templates/postform.html", p)
}

func DisplayPostPreview(w http.ResponseWriter, r *http.Request, p *PostForm) {

	moodState, _ := strconv.Atoi(p.Fields["mood"])
	post := socialmedia.NewPost("Anonymous Gopher", socialmedia.MoodState(moodState), p.Fields["caption"], p.Fields["messageBody"], "", "", "", nil)
	RenderTemplate(w, WebAppRoot+"/templates/socialmediapost.html", post)
}

func PopulatePostFormFields(r *http.Request, p *PostForm) {

	for _, fieldName := range p.FieldNames {
		p.Fields[fieldName] = r.FormValue(fieldName)
	}

}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidatePostForm(w http.ResponseWriter, r *http.Request, p *PostForm) {

	p.CurrentMood, _ = strconv.Atoi(r.FormValue("mood"))

	PopulatePostFormFields(r, p)
	// Check if caption was filled out
	if r.FormValue("caption") == "" {
		p.Errors["captionError"] = "The caption field is required."
	}

	// Check if first name was filled out
	if r.FormValue("messageBody") == "" {
		p.Errors["messageBodyError"] = "The post message body is required."
	}

	if len(p.Errors) > 0 {
		DisplayPostForm(w, r, p)
	} else {
		DisplayPostPreview(w, r, p)
	}

}

func PostPreviewHandler(w http.ResponseWriter, r *http.Request) {

	p := PostForm{}
	p.FieldNames = []string{"caption", "messageBody", "mood"}
	p.Fields = make(map[string]string)
	p.Errors = make(map[string]string)
	p.Moods = MoodStates
	p.MoodEmoji = MoodEmoji
	p.CurrentMood = 0

	switch r.Method {

	case "GET":
		DisplayPostForm(w, r, &p)
	case "POST":
		ValidatePostForm(w, r, &p)
	default:
		DisplayPostPreview(w, r, &p)
	}

}

func init() {

	MoodStates = make(map[string]int)
	MoodStates["Neutral"] = int(socialmedia.MoodStateNeutral)
	MoodStates["Happy"] = int(socialmedia.MoodStateHappy)
	MoodStates["Sad"] = int(socialmedia.MoodStateSad)
	MoodStates["Angry"] = int(socialmedia.MoodStateAngry)
	MoodStates["Hopeful"] = int(socialmedia.MoodStateHopeful)
	MoodStates["Thrilled"] = int(socialmedia.MoodStateThrilled)
	MoodStates["Bored"] = int(socialmedia.MoodStateBored)
	MoodStates["Shy"] = int(socialmedia.MoodStateShy)
	MoodStates["Comical"] = int(socialmedia.MoodStateComical)
	MoodStates["On Cloud Nine"] = int(socialmedia.MoodStateOnCloudNine)

	MoodEmoji = make(map[string]string)
	MoodEmoji["Neutral"] = socialmedia.MoodsEmoji[socialmedia.MoodStateNeutral]
	MoodEmoji["Happy"] = socialmedia.MoodsEmoji[socialmedia.MoodStateHappy]
	MoodEmoji["Sad"] = socialmedia.MoodsEmoji[socialmedia.MoodStateSad]
	MoodEmoji["Angry"] = socialmedia.MoodsEmoji[socialmedia.MoodStateAngry]
	MoodEmoji["Hopeful"] = socialmedia.MoodsEmoji[socialmedia.MoodStateHopeful]
	MoodEmoji["Thrilled"] = socialmedia.MoodsEmoji[socialmedia.MoodStateThrilled]
	MoodEmoji["Bored"] = socialmedia.MoodsEmoji[socialmedia.MoodStateBored]
	MoodEmoji["Shy"] = socialmedia.MoodsEmoji[socialmedia.MoodStateShy]
	MoodEmoji["Comical"] = socialmedia.MoodsEmoji[socialmedia.MoodStateComical]
	MoodEmoji["On Cloud Nine"] = socialmedia.MoodsEmoji[socialmedia.MoodStateOnCloudNine]

}

package forms

import (
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models/socialmedia"

	"go.isomorphicgo.org/go/isokit"
)

type SocialMediaPostForm struct {
	isokit.BasicForm
	Moods       map[string]int
	MoodEmoji   map[string]string
	CurrentMood int
	PageTitle   string
}

var MoodStates map[string]int
var MoodEmoji map[string]string

func NewSocialMediaPostForm(formParams *isokit.FormParams) *SocialMediaPostForm {
	prefillFields := []string{"caption", "messageBody"}
	fields := make(map[string]string)
	errors := make(map[string]string)
	s := &SocialMediaPostForm{}
	s.SetPrefillFields(prefillFields)
	s.SetFields(fields)
	s.SetErrors(errors)
	s.SetFormParams(formParams)
	s.Moods = MoodStates
	s.MoodEmoji = MoodEmoji
	s.CurrentMood = 0
	return s
}

func (s *SocialMediaPostForm) Validate() bool {

	s.RegenerateErrors()
	s.PopulateFields()

	// Check if the capition field was filled out
	if isokit.FormValue(s.FormParams(), "caption") == "" {
		s.SetError("caption", "The caption field is required.")
	}

	// Check if the message body was filled out
	if isokit.FormValue(s.FormParams(), "messageBody") == "" {
		s.SetError("messageBody", "You have to fill out the post body field!")
	}

	if len(s.Errors()) > 0 {
		return false

	} else {
		return true
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

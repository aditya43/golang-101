// Package SocialMedia implements common functionality needed for social media web applications.
package socialmedia

import (
	"time"
)

//go:generate stringer -type=MoodState
type MoodState int

// All possible mood states.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

// AuditableContent types are meant to be embeded into types we want to keep a
// check on for auditing purposes
type AuditableContent struct {
	TimeCreated        time.Time `json:"timeCreated"`
	TimeModified       time.Time `json:"timeModified"`
	TimeCreatedUnixTS  int       `json:"timeCreatedUnixTS"`
	TimeModifiedUnixTS int       `json:"timeModifiedUnixTS"`
	CreatedBy          string    `json:"createdBy"`
	ModifiedBy         string    `json:"modifiedBy"`
}

// Post represents a Social Media Post type.
type Post struct {
	AuditableContent           // Embedded type
	UUID             string    `json:"uuid"`
	RawMoodValue     int       `json:"rawMoodValue"`
	Username         string    `json:"username"`
	ProfileImagePath string    `json:"profileImagePath"`
	Caption          string    `json:"caption"`
	MessageBody      string    `json:"messageBody"`
	URL              string    `json:"url"`
	ImageURI         string    `json:"imageURI"`
	ThumbnailURI     string    `json:"thumbnailURI"`
	Keywords         []string  `json:"keywords"`
	Likers           []string  `json:"likers"`
	AuthorMood       MoodState `json:"authorMood"`
	AuthorMoodEmoji  string    `json:"authorMoodEmoji"`
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood states.
var Moods map[string]MoodState
var MoodsEmoji map[MoodState]string

// The init() function is responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}

	MoodsEmoji = map[MoodState]string{MoodStateNeutral: "\xF0\x9F\x98\x90", MoodStateHappy: "\xF0\x9F\x98\x8A", MoodStateSad: "\xF0\x9F\x98\x9E", MoodStateAngry: "\xF0\x9F\x98\xA0", MoodStateHopeful: "\xF0\x9F\x98\x8C", MoodStateThrilled: "\xF0\x9F\x98\x81", MoodStateBored: "\xF0\x9F\x98\xB4", MoodStateShy: "\xF0\x9F\x98\xB3", MoodStateComical: "\xF0\x9F\x98\x9C", MoodStateOnCloudNine: "\xF0\x9F\x98\x82"}
}

// NewPost is responsible for creating an instance of the Post type.
func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {
	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageURI, ThumbnailURI: thumbnailURI, AuthorMood: mood, Keywords: keywords, AuditableContent: auditableContent, AuthorMoodEmoji: MoodsEmoji[mood]}
}

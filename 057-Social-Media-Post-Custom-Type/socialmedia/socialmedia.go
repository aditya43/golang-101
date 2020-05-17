package socialmedia

import (
	"time"
)

//go:generate stringer -type=MoodState
type MoodState int

// Here we define all the possible mood states using an
// iota enumerator.
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

// This is a type we embed into types we want to keep a
// check on for auditing purposes
type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// This is the type that represents a Social Media Post
type Post struct {
	AuditableContent // Embedded type
	Caption          string
	MessageBody      string
	URL              string
	ImageURI         string
	ThumbnailURI     string
	Keywords         []string
	Likers           []string
	AuthorMood       MoodState
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood state
var Moods map[string]MoodState

// The init() function is responsible for initializing the mood state
func init() {
	Moods = map[string]MoodState{
		"netural":   MoodStateNeutral,
		"happy":     MoodStateHappy,
		"sad":       MoodStateSad,
		"angry":     MoodStateAngry,
		"hopeful":   MoodStateHopeful,
		"thrilled":  MoodStateThrilled,
		"bored":     MoodStateBored,
		"shy":       MoodStateShy,
		"comical":   MoodStateComical,
		"cloudnine": MoodStateOnCloudNine,
	}
}

// Constructor for "Post" struct: This is the function responsible for creating a new social media post.
func NewPost(
	username string,
	mood MoodState,
	caption string,
	messageBody string,
	url string,
	imageURI string,
	thumbnailURI string,
	keywords []string) *Post {

	auditableContent := AuditableContent{
		CreatedBy:   username,
		TimeCreated: time.Now(),
	}

	return &Post{
		AuditableContent: auditableContent,
		Caption:          caption,
		MessageBody:      messageBody,
		URL:              url,
		ImageURI:         imageURI,
		ThumbnailURI:     thumbnailURI,
		AuthorMood:       mood,
		Keywords:         keywords,
	}
}

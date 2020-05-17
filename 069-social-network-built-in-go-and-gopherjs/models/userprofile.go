package models

import "github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/forms"

type UserProfile struct {
	PageTitle        string `json:"pageTitle" bson:"pageTitle"`
	UUID             string `json:"uuid" bson:"uuid"`
	Username         string `json:"username" bson:"username"`
	About            string `json:"about" bson:"about"`
	Location         string `json:"location" bson:"location"`
	Interests        string `json:"interests" bson:"interests"`
	ProfileImagePath string `json:"profileImagePath" bson:"profileImagePath"`
	Form             *forms.MyProfileForm
}

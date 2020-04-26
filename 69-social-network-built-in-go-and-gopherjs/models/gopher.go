package models

type Gopher struct {
	UUID             string `json:"uuid" bson:"uuid"`
	Username         string `json:"username" bson:"username"`
	FirstName        string `json:"firstName" bson:"firstName"`
	LastName         string `json:"lastName" bson:"lastName"`
	Location         string `json:"location" bson:"location"`
	ProfileImagePath string `json:"profileImagePath" bson:"profileImagePath"`
}

package datastore

import (
	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models/socialmedia"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"
)

type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
	GetUserProfile(uuid string) (*models.UserProfile, error)
	UpdateUserProfile(uuid, about, location, interests string) error
	UpdateUserProfileImage(uuid, profileImagePath string) error
	FindGophers(owner string, searchTerm string) ([]models.Gopher, error)
	FriendsList(owner string) ([]models.Gopher, error)
	FollowGopher(owner string, friend string) error
	UnfollowGopher(owner string, friend string) error
	SavePost(owner string, title string, body string, mood int) error
	FetchPosts(owner string) ([]socialmedia.Post, error)
	GetGopherProfile(username string) (*models.UserProfile, error)
}

const (
	MYSQL = iota
	//MONGODB
	//REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
		/*
			case MONGODB:
				return NewMongoDBDatastore(dbConnectionString)
			case REDIS:
				return NewRedisDatastore(dbConnectionString)
		*/
	}

	return nil, nil
}

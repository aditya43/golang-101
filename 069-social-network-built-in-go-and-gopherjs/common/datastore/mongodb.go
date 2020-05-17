package datastore

import (
	"fmt"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoDBDatastore struct {
	*mgo.Session
}

func NewMongoDBDatastore(url string) (*MongoDBDatastore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDBDatastore{
		Session: session,
	}, nil
}

func (m *MongoDBDatastore) CreateUser(user *models.User) error {

	session := m.Copy()

	defer session.Close()
	userCollection := session.DB("gopherface").C("User")
	err := userCollection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDBDatastore) GetUser(username string) (*models.User, error) {

	session := m.Copy()
	defer session.Close()
	userCollection := session.DB("gopherface").C("User")
	u := models.User{}
	err := userCollection.Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}
	fmt.Println("created")
	return &u, nil

}

func (m *MongoDBDatastore) Close() {
	m.Close()
}

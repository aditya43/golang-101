package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"

	"github.com/mediocregopher/radix.v2/pool"
)

type RedisDatastore struct {
	*pool.Pool
}

func NewRedisDatastore(address string) (*RedisDatastore, error) {

	connectionPool, err := pool.New("tcp", address, 10)
	if err != nil {
		return nil, err
	}
	return &RedisDatastore{
		Pool: connectionPool,
	}, nil
}

func (r *RedisDatastore) CreateUser(user *models.User) error {

	userJSON, err := json.Marshal(*user)
	if err != nil {
		return err
	}

	if r.Cmd("SET", "user:"+user.Username, string(userJSON)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

func (r *RedisDatastore) GetUser(username string) (*models.User, error) {

	exists, err := r.Cmd("EXISTS", "user:"+username).Int()

	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, nil
	}

	var u models.User

	userJSON, err := r.Cmd("GET", "user:"+username).Str()

	fmt.Println("userJSON: ", userJSON)

	if err != nil {
		log.Print(err)

		return nil, err
	}

	if err := json.Unmarshal([]byte(userJSON), &u); err != nil {
		log.Print(err)
		return nil, err
	}

	return &u, nil
}

func (r *RedisDatastore) Close() {

	r.Close()
}

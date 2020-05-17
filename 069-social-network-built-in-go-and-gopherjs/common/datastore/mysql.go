package datastore

import (
	"database/sql"
	"log"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models/socialmedia"

	"github.com/aditya43/golang/69-social-network-built-in-go-and-gopherjs/models"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatastore struct {
	*sql.DB
}

func NewMySQLDatastore(dataSourceName string) (*MySQLDatastore, error) {

	connection, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &MySQLDatastore{
		DB: connection}, nil
}

func (m *MySQLDatastore) CreateUser(user *models.User) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.UUID, user.Username, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLDatastore) GetUser(username string) (*models.User, error) {

	stmt, err := m.Prepare("SELECT uuid, username, first_name, last_name, email, password_hash, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM user WHERE username = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(username)
	u := models.User{}
	err = row.Scan(&u.UUID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}

func (m *MySQLDatastore) Close() {
	m.Close()
}

func (m *MySQLDatastore) GetUserProfile(uuid string) (*models.UserProfile, error) {

	stmt, err := m.Prepare("SELECT uuid, about, location, interests, profile_image_path FROM user_profile WHERE uuid = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(uuid)
	u := models.UserProfile{}
	err = row.Scan(&u.UUID, &u.About, &u.Location, &u.Interests, &u.ProfileImagePath)
	return &u, err
}

func (m *MySQLDatastore) GetGopherProfile(username string) (*models.UserProfile, error) {

	stmt, err := m.Prepare("SELECT up.uuid, up.about, up.location, up.interests, up.profile_image_path, u.username FROM user_profile up, user u WHERE u.uuid = up.uuid and u.username = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(username)
	u := models.UserProfile{}
	err = row.Scan(&u.UUID, &u.About, &u.Location, &u.Interests, &u.ProfileImagePath, &u.Username)
	return &u, err
}

func (m *MySQLDatastore) UpdateUserProfile(uuid, about, location, interests string) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user_profile(uuid, about, location, interests) VALUES(?, ?, ?, ?) ON DUPLICATE KEY UPDATE about=?, location=?, interests=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, about, location, interests, about, location, interests)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (m *MySQLDatastore) UpdateUserProfileImage(uuid, profileImagePath string) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user_profile(uuid, profile_image_path) VALUES(?, ?) ON DUPLICATE KEY UPDATE profile_image_path=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, profileImagePath, profileImagePath)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (m *MySQLDatastore) FindGophers(owner string, searchTerm string) ([]models.Gopher, error) {

	searchTerm = "%" + searchTerm + "%"
	gophers := make([]models.Gopher, 0)
	stmt, err := m.Prepare("SELECT u.uuid, u.username, u.first_name, u.last_name, up.location, up.profile_image_path FROM user u, user_profile up WHERE u.uuid = up.uuid AND CONCAT(u.first_name, '', u.last_name) LIKE ? and u.uuid <> ? and u.uuid not in (select friend_uuid from friend_relation where owner_uuid = ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(searchTerm, owner, owner)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		g := models.Gopher{}
		err := rows.Scan(&g.UUID, &g.Username, &g.FirstName, &g.LastName, &g.Location, &g.ProfileImagePath)
		if err != nil {
			return nil, err
		}
		gophers = append(gophers, g)
	}
	return gophers, nil

}

func (m *MySQLDatastore) FriendsList(owner string) ([]models.Gopher, error) {

	gophers := make([]models.Gopher, 0)
	stmt, err := m.Prepare("SELECT u.uuid, u.username, u.first_name, u.last_name, up.location, up.profile_image_path FROM user u, user_profile up, friend_relation f WHERE u.uuid = up.uuid AND u.uuid = f.friend_uuid AND f.owner_uuid=?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(owner)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		g := models.Gopher{}
		err := rows.Scan(&g.UUID, &g.Username, &g.FirstName, &g.LastName, &g.Location, &g.ProfileImagePath)
		if err != nil {
			return nil, err
		}
		gophers = append(gophers, g)
	}
	return gophers, nil

}

func (m *MySQLDatastore) FollowGopher(owner string, friend string) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO friend_relation(owner_uuid, friend_uuid) VALUES(?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(owner, friend)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (m *MySQLDatastore) UnfollowGopher(owner string, friend string) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("DELETE FROM friend_relation WHERE owner_uuid = ? and friend_uuid = ? LIMIT 1")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(owner, friend)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (m *MySQLDatastore) SavePost(owner string, title string, body string, mood int) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO post(uuid, title, body, mood) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(owner, title, body, mood)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func (m *MySQLDatastore) FetchPosts(owner string) ([]socialmedia.Post, error) {

	posts := make([]socialmedia.Post, 0)
	stmt, err := m.Prepare("select p.uuid, p.title, p.body, p.mood, UNIX_TIMESTAMP(p.created_ts), UNIX_TIMESTAMP(p.updated_ts), up.profile_image_path, u.username from post p, user u, user_profile up where p.uuid = u.uuid and p.uuid = up.uuid and (p.uuid = ? or p.uuid in (select friend_uuid from friend_relation where owner_uuid=?) ) order by p.created_ts desc")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(owner, owner)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := socialmedia.Post{}
		err := rows.Scan(&p.UUID, &p.Caption, &p.MessageBody, &p.RawMoodValue, &p.TimeCreatedUnixTS, &p.TimeModifiedUnixTS, &p.ProfileImagePath, &p.Username)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil

}

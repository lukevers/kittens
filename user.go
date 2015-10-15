/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID          int
	Username    string `sql:"unique"`
	Password    string
	Email       string `sql:"unique"`
	Twofa       bool
	TwofaSecret string
	Bots        []*Bot
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetUser(by, value interface{}) *User {
	var user User
	db.Where(fmt.Sprintf("%s = ?", by), value).First(&user)
	return &user
}

func (u User) Related() *User {
	db.Model(&u).Association("Bots").Find(&u.Bots)
	for b := range u.Bots {
		db.Model(&u.Bots[b]).Association("Channels").Find(&u.Bots[b].Channels)
		for c := range u.Bots[b].Channels {
			db.Model(&u.Bots[b].Channels[c]).Association("Plugins").Find(&u.Bots[b].Channels[c].Plugins)
		}
	}

	return &u
}

func (u User) AttemptPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u User) SetPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Could not generate hash from password: ", err)
	}

	u.Password = string(hash)
	db.Save(&u)
}

func (u User) SetUsername(username string) error {
	// Check to see if the username is the same -- if so, don't update.
	if u.Username == username {
		return errors.New("No difference in username")
	}

	// Check to see if another user already has this username.
	user := GetUser("username", username)
	if user.ID != 0 {
		return errors.New("Username already exists")
	}

	// If we got this far, let's just update the username.
	u.Username = username
	db.Save(&u)

	return nil
}

func (u User) SetEmail(email string) error {
	// Check to see if the email is the same -- if so, don't update.
	if u.Email == email {
		return errors.New("No difference in email")
	}

	// Check to see if another user already has this email.
	user := GetUser("email", email)
	if user.ID != 0 {
		return errors.New("Email already exists")
	}

	// If we got this far, let's just update the email.
	u.Email = email
	db.Save(&u)

	return nil
}

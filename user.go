package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	Id          uint
	Username    string `sql:"unique"`
	Password    string
	Twofa       bool
	TwofaSecret string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func GetUser(by, value interface{}) *User {
	var user User
	db.Where(fmt.Sprintf("%s = ?", by), value).First(&user)
	return &user
}

func (u User) SetPassword(password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Could not generate hash from password: ", err)
	}

	u.Password = string(hash)
	db.Save(&u)
}

func (u User) AttemptPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

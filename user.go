package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Session store for users
var store = sessions.NewFilesystemStore(
	// Path
	"sessions",
	// Secret key with strength set to 64
	[]byte(securecookie.GenerateRandomKey(64)),
)

type User struct {
	// ID is a uint32 that is a users identification number.
	ID uint32

	// Username is the name that a user uses in order to sign in.
	Username string

	// Password is the secret key that a user types in along
	// with a matching username in order to sign in.
	Password string
}

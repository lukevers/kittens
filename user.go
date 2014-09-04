package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"time"
)

// Session store for users
var store = sessions.NewFilesystemStore(
	// Path
	"app/sessions",
	// Secret key with strength set to 64
	[]byte(securecookie.GenerateRandomKey(64)),
)

type User struct {
	// ID is an int64 that is a users identification
	// number.
	Id uint64

	// Username is a string with max-size set to 255
	// and is the username that a user will use when
	// logging in to the web interface.
	Username string `sql:"size:255;unique"`

	// Password is a string with max-size set to 255
	// and is the password that a user will use when
	// logging in to the web interface.
	Password string `sql:"size:255"`

	// CreatedAt is a timestamp of when the specific
	// user was created at.
	CreatedAt time.Time

	// UpdatedAt is a timestamp of when the specific
	// user was last updated at.
	UpdatedAt time.Time
}

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

	// Admin is a bool that specifies if the current
	// user is an administrator or not.
	Admin bool

	// CreatedAt is a timestamp of when the specific
	// user was created at.
	CreatedAt time.Time

	// UpdatedAt is a timestamp of when the specific
	// user was last updated at.
	UpdatedAt time.Time

	// Servers is a slice of Server structs that define
	// what servers the user owns
	Servers []*Server `sql:"-"`
}

// Owns Server checks if the given *Server is owned by
// the current user given.
func (user *User) OwnsServer(server *Server) bool {
	// If we go to a server that doesn't exist, we
	// automatically know we don't own that server
	if server == nil || user == nil {
		return false
	}

	// Check
	for _, s := range user.Servers {
		if s.Id == server.Id {
			return true
		}
	}

	// Return false if we never found it
	return false
}

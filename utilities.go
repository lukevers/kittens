package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"errors"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

// Get Server From Request takes a *http.Request in and returns
// the server ID that is being looked at. It returns an error if
// it can't find it.
func GetServerFromRequest(req *http.Request) (*Server, error) {
	// Figure out what {id} is in "/server/{id}"
	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 16)
	if err != nil {
		warnf("Error converting server id: %s", err)
	}

	// Get our server from our slice of servers
	for _, s := range servers {
		if s.Id == uint64(id) {
			return s, nil
		}
	}

	return nil, errors.New("Could not find server")
}

// Get Channel From Request takes a *Server and *http.Request
// and returns the *Channel that is being looked at. It returns
// and error if it can't find it.
func GetChannelFromRequest(s *Server, req *http.Request) (*Channel, error) {
	// Figure out what {channel} is
	channel := mux.Vars(req)["channel"]

	// Get our channel from our slice of channels
	for _, c := range s.Channels {
		if channel == c.Name {
			return c, nil
		}
	}

	return nil, errors.New("Could not find channel")
}

// Is Logged In checks if the user has a session or not.
// If the user does not have a session that matches with
// what we have, then the user is not logged in.
func IsLoggedIn(req *http.Request) bool {
	// Check for session
	session, _ := store.Get(req, "user")
	return *noAuthFlag || !session.IsNew
}

// WhoAmI figures out who exactly is using the current
// session (what user is), and it returns the *User from
// the slice of Users that we have.
func WhoAmI(req *http.Request) *User {
	session, _ := store.Get(req, "user")
	for _, user := range users {
		if session.Values["username"] == user.Username {
			return user
		}
	}

	return nil
}

// Hash Password takes a string and hashes that password
// and returns it as a string. It handles errors that are
// returned from bcrypt.GenerateFromPassword, and is a
// wrapper around having to use []byte everywhere.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		warnf("Error hashing password: %s", err)
	}

	return string(hash)
}

// Password Matches Hash takes a plaintext password and uses
// bcrypt.CompareHashAndPassword to check against the hashed
// password we're checking against from the database. The
// func from bcrypt returns nil if the passwords match, and
// an error otherwise, so we're checking if bcrypt's func
// returns nil or not and that's how we're determining if the
// hashes match or not.
func PasswordMatchesHash(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// Refresh Templates recompiles the templates. We use this a lot,
// so it's better to have it in once place than in 20 places.
func RefreshTemplates(req *http.Request) *template.Template {
	return template.Must(template.New("").Funcs(AddTemplateFunctions(req)).ParseGlob("app/views/*"))
}

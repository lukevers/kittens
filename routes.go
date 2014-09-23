package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Route is the gorilla/mux router we use
var route *mux.Router

func AddRoutes() {
	// Handles GET requests for "/" which is our root page.
	route.HandleFunc("/", HandleRoot)

	// Handles GET requests to "/login" which displays a form that a user
	// can use to try and login.
	route.HandleFunc("/login", HandleLogin).Methods("GET")

	// Handles POST requests for "/login" which tests if a user is
	// logging in with correct details or not.
	route.HandleFunc("/login", HandleLoginForm).Methods("POST")

	// Handles GET requests to "/login/2fa" which displays a form that a
	// user can use to try their 2fa token on.
	route.HandleFunc("/login/2fa", HandleLogin2FA).Methods("GET")

	// Handles POST requests to "/login/2fa" which tests if a users 2fa
	// token is correct or not.
	route.HandleFunc("/login/2fa", HandleLoginForm2FA).Methods("POST")

	// Handle logout requests which removes the session and logs the user out
	route.HandleFunc("/logout", HandleLogout)

	// Handles GET requests for "/settings" which is a page where users
	// can update their settings.
	route.HandleFunc("/settings", HandleSettings).Methods("GET")

	// Handles POST requests for "/settings" which is a page where users
	// can update their settings. POSTing here will update settings.
	route.HandleFunc("/settings", HandleUpdateSettings).Methods("POST")

	// Handles GET requests for "/settings/2fa/generate" which is a page
	// that generates a QR code for Two Factor Auth.
	route.HandleFunc("/settings/2fa/generate", HandleGenerate2FA).Methods("GET")

	// Handles POST requests for "/settings/2fa/verify" which checks to
	// see if a 2FA token is correct. If it is, then we remove the temp
	// secret key from the session and add it to the database.
	route.HandleFunc("/settings/2fa/verify", HandleVerify2FA).Methods("POST")

	// Handles POST requests for "/settings/2fa/disable" which disables
	// 2FA for the users account.
	route.HandleFunc("/settings/2fa/disable", HandleDisable2FA).Methods("POST")

	// Handles GET requests for "/users" which is an admin-only page
	route.HandleFunc("/users", HandleUsers).Methods("GET")

	// Handles POST requests for "/users/new" which is a form where
	// new users can be added.
	route.HandleFunc("/users/new", HandleNewUser).Methods("POST")

	// Handles POST requests for "/users/delete" which is how users
	// can be deleted.
	route.HandleFunc("/users/delete", HandleUserDelete).Methods("POST")

	// Handles POST requests for "/users/admin" which is a form where
	// administrators can promote/demote users.
	route.HandleFunc("/users/admin", HandleUserAdminSwitch).Methods("POST")

	// Handles GET requests for "/server/new" which is a page where a
	// user can add a new server.
	route.HandleFunc("/server/new", HandleNew).Methods("GET")

	// Handles POST requests for "/server/new" which adds a new server
	route.HandleFunc("/server/new", HandleAddNew).Methods("POST")

	// Handles GET requests for "/server/{id}" which is a server page
	route.HandleFunc("/server/{id}", HandleServer).Methods("GET")

	// Handles POST requests for "/server/{id}" which is an endpoint where
	// server information can be updated.
	route.HandleFunc("/server/{id}", HandleUpdateServer).Methods("POST")

	// Handles GET requests for "/server/{id}/channel/" which is a page
	// that looks for a URL fragment at the end of the url. JavaScript
	// then takes that URL fragment and URL encodes the fragment. After
	// URL encoding the fragment we're redirected the correct channel page.
	route.HandleFunc("/server/{id}/channel/", HandleChannelRedirect).Methods("GET")

	// Handles GET requests for "/server/{id}/channel/{channel}" which
	// is a page for a specific channel for a specific server.
	route.HandleFunc("/server/{id}/channel/{channel}", HandleChannel).Methods("GET")

	// Handles POST requests for "/server/{id}/enable" which takes a bool
	// and enables--if it is disabled--the server if the bool is true, and
	// disables--if it is enabled--the server if the bool is false.
	route.HandleFunc("/server/{id}/enable", HandleEnableServer).Methods("POST")

	// Handles POST requests for "/server/{id}/channel/join" which takes
	// a specific channel and joins it.
	route.HandleFunc("/server/{id}/channel/join", HandleJoinChannel).Methods("POST")

	// Handles POST requests for "/server/{id}/channel/part" which takes
	// a specific channel and parts it.
	route.HandleFunc("/server/{id}/channel/part", HandlePartChannel).Methods("POST")

	// Handle all other static files and folders (eg. CSS/JS).
	route.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
}

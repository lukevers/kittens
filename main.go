package main

import (
	"flag"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var (
	err     error
	wg      sync.WaitGroup
	users   []*User
	servers []*Server
)

var (
	templates = template.Must(template.New("").Funcs(AddTemplateFunctions(nil)).ParseGlob("app/views/*"))
)

func main() {
	// Parse flags
	flag.Parse()

	// Initialize database
	InitDatabase()

	// Remove old sessions
	CleanSessions()

	info("Starting webserver")

	// Web server
	r := mux.NewRouter()

	// Handles GET requests for "/" which is our root page.
	r.HandleFunc("/", HandleRoot)

	// Handles GET requests to "/login" which displays a form that a user
	// can use to try and login.
	r.HandleFunc("/login", HandleLogin).Methods("GET")

	// Handles POST requests for "/login" which tests if a user is
	// logging in with correct details or not.
	r.HandleFunc("/login", HandleLoginForm).Methods("POST")

	// Handles GET requests to "/login/2fa" which displays a form that a
	// user can use to try their 2fa token on.
	r.HandleFunc("/login/2fa", HandleLogin2FA).Methods("GET")

	// Handles POST requests to "/login/2fa" which tests if a users 2fa
	// token is correct or not.
	r.HandleFunc("/login/2fa", HandleLoginForm2FA).Methods("POST")

	// Handle logout requests which removes the session and logs the user out
	r.HandleFunc("/logout", HandleLogout)

	// Handles GET requests for "/settings" which is a page where users
	// can update their settings.
	r.HandleFunc("/settings", HandleSettings).Methods("GET")

	// Handles POST requests for "/settings" which is a page where users
	// can update their settings. POSTing here will update settings.
	r.HandleFunc("/settings", HandleUpdateSettings).Methods("POST")

	// Handles GET requests for "/settings/2fa/generate" which is a page
	// that generates a QR code for Two Factor Auth.
	r.HandleFunc("/settings/2fa/generate", HandleGenerate2FA).Methods("GET")

	// Handles POST requests for "/settings/2fa/verify" which checks to
	// see if a 2FA token is correct. If it is, then we remove the temp
	// secret key from the session and add it to the database.
	r.HandleFunc("/settings/2fa/verify", HandleVerify2FA).Methods("POST")

	// Handles POST requests for "/settings/2fa/disable" which disables
	// 2FA for the users account.
	r.HandleFunc("/settings/2fa/disable", HandleDisable2FA).Methods("POST")

	// Handles GET requests for "/users" which is an admin-only page
	r.HandleFunc("/users", HandleUsers).Methods("GET")

	// Handles POST requests for "/users/new" which is a form where
	// new users can be added.
	r.HandleFunc("/users/new", HandleNewUser).Methods("POST")

	// Handles POST requests for "/users/delete" which is how users
	// can be deleted.
	r.HandleFunc("/users/delete", HandleUserDelete).Methods("POST")

	// Handles POST requests for "/users/admin" which is a form where
	// administrators can promote/demote users.
	r.HandleFunc("/users/admin", HandleUserAdminSwitch).Methods("POST")

	// Handles GET requests for "/server/new" which is a page where a
	// user can add a new server.
	r.HandleFunc("/server/new", HandleNew).Methods("GET")

	// Handles POST requests for "/server/new" which adds a new server
	r.HandleFunc("/server/new", HandleAddNew).Methods("POST")

	// Handles GET requests for "/server/{id}" which is a server page
	r.HandleFunc("/server/{id}", HandleServer).Methods("GET")

	// Handles POST requests for "/server/{id}" which is an endpoint where
	// server information can be updated.
	r.HandleFunc("/server/{id}", HandleUpdateServer).Methods("POST")

	// Handles GET requests for "/server/{id}/channel/" which is a page
	// that looks for a URL fragment at the end of the url. JavaScript
	// then takes that URL fragment and URL encodes the fragment. After
	// URL encoding the fragment we're redirected the correct channel page.
	r.HandleFunc("/server/{id}/channel/", HandleChannelRedirect).Methods("GET")

	// Handles GET requests for "/server/{id}/channel/{channel}" which
	// is a page for a specific channel for a specific server.
	r.HandleFunc("/server/{id}/channel/{channel}", HandleChannel).Methods("GET")

	// Handles POST requests for "/server/{id}/enable" which takes a bool
	// and enables--if it is disabled--the server if the bool is true, and
	// disables--if it is enabled--the server if the bool is false.
	r.HandleFunc("/server/{id}/enable", HandleEnableServer).Methods("POST")

	// Handles POST requests for "/server/{id}/channel/join" which takes
	// a specific channel and joins it.
	r.HandleFunc("/server/{id}/channel/join", HandleJoinChannel).Methods("POST")

	// Handles POST requests for "/server/{id}/channel/part" which takes
	// a specific channel and parts it.
	r.HandleFunc("/server/{id}/channel/part", HandlePartChannel).Methods("POST")

	// Handle all other static files and folders (eg. CSS/JS).
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	// Get all users
	db.Find(&users, &User{})
	for _, user := range users {
		// Get all of the servers for each user
		db.Table("servers").Where("user_id = ?", user.Id).Find(&user.Servers)

		// Create servers
		for _, server := range user.Servers {
			// Get all of the channels for each server
			db.Table("channels").Where("server_id = ?", server.Id).Find(&server.Channels)

			// Add to slice of server
			servers = append(servers, server)

			// Add to wait group
			wg.Add(1)

			// Start our goroutine
			go server.Create()
		}
	}

	http.Handle("/", r)
	http.ListenAndServe(*interfaceFlag+":"+strconv.Itoa(*portFlag), nil)
	infof("Webserver running on port %s", *portFlag)

	wg.Wait()
}

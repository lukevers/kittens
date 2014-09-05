package main

import (
	"flag"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
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

	// Load the configuration file
	config, err = ReadConfig(*configPathFlag)
	if err != nil {
		warn("Could not load configuration file.")
		warnf("Error: %s", err)
		warn("Exiting with exit status 1")
		os.Exit(1)
	}

	// Update config values from flags
	config.UpdateFromFlags()

	// Initialize database
	InitDatabase()

	verb("Loaded configuration file")
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

	// Handle logout requests which removes the session and logs the user
	// out.
	r.HandleFunc("/logout", HandleLogout)

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
	}

	// Get all servers
	db.Find(&servers, &Server{})

	// Create servers
	for _, s := range servers {
		// Get all of the channels for each server
		db.Table("channels").Where("server_id = ?", s.Id).Find(&s.Channels)

		// Wait group
		wg.Add(1)

		// Start our goroutine
		go s.CreateAndConnect(true)
	}

	http.Handle("/", r)
	http.ListenAndServe(config.Interface+":"+strconv.Itoa(config.Port), nil)
	infof("Webserver running on port %s", config.Port)

	wg.Wait()
}

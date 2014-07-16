package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	config  *Config
	err     error
	wg      sync.WaitGroup
	clients []*Server
)

var (
	templates        = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	nextID    uint16 = 0
)

func main() {
	// Load the configuration file
	config, err = ReadConfig("config.json")

	if err != nil {
		warn("Could not load configuration file.")
		warnf("Error: %s", err)
		warn("Exiting with exit status 1")
		os.Exit(1)
	}

	verb("Loaded configuration file")
	info("Starting webserver")

	// Web server
	r := mux.NewRouter()

	// Handles GET requests for "/" which is our root page.
	r.HandleFunc("/", HandleRoot)

	// Handles GET requests for "/server/{id}" which is a server page
	r.HandleFunc("/server/{id}", HandleServer).Methods("GET")

	// Handles POST requests for "/server/{id}" which is an endpoint where
	// server information can be updated.
	r.HandleFunc("/server/{id}", HandleUpdateServer).Methods("POST")

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

	info("Beginning to create bots")

	for _, s := range config.Servers {
		wg.Add(1)
		s.ID = nextID
		nextID++
		go s.CreateAndConnect(true)
	}

	verbf("port %s", config.Port)

	// Check if config.Port exists
	if config.Port == 0 {
		// If it does not exist, let's just give it 3000
		config.Port = 3000
	}

	// Check if config.Interface exists
	if config.Interface == "" {
		// If it does not exist let's give it 0.0.0.0
		config.Interface = "0.0.0.0"
	}

	http.Handle("/", r)
	http.ListenAndServe(config.Interface+":"+strconv.Itoa(config.Port), nil)
	infof("Webserver running on port %s", config.Port)

	wg.Wait()
}

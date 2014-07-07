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
	r.HandleFunc("/", HandleRoot)

	// Handle /server/{id}
	r.HandleFunc("/server/{id}", HandleServer).Methods("GET")
	r.HandleFunc("/server/{id}", HandleUpdateServer).Methods("POST")
	r.HandleFunc("/server/{id}/enable", HandleEnableServer).Methods("POST")
	r.HandleFunc("/server/{id}/channel/join", HandleJoinChannel).Methods("POST")
	r.HandleFunc("/server/{id}/channel/part", HandlePartChannel).Methods("POST")
	r.HandleFunc("/server/{id}/channel/{channel}", HandleChannel).Methods("GET")

	// Handle static
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

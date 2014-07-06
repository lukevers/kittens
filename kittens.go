package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"sync"
)

var (
	config   *Config
	err      error
	wg       sync.WaitGroup
	clients  []*Server
)

var (
	templates     = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	nextID uint16 = 0
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
	r.HandleFunc("/server/{id}", UpdateServer).Methods("POST")
	r.HandleFunc("/server/{id}/enable", EnableServer).Methods("POST")
	r.HandleFunc("/server/{id}/channel/join", JoinChannel).Methods("POST")
	r.HandleFunc("/server/{id}/channel/part", PartChannel).Methods("POST")

	// Handle static
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	info("Beginning to create bots")

	for _, s := range config.Servers {
		wg.Add(1)
		s.ID = nextID; nextID++
		go s.CreateAndConnect(true)
	}

	http.Handle("/", r)
	http.ListenAndServe(config.Port, nil)
	infof("Webserver running on port %s", config.Port)

	wg.Wait()
}

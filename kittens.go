package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"sync"
)

var (
	config *Config
	err    error
	wg     sync.WaitGroup
	cli    []*Server
)

var (
	templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
)

func main() {

	// Load the configuration file
	config, err = ReadConfig("example.config.json")

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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))

	// Template functions

	// Back to bots
	info("Beginning to create bots")

	for _, s := range config.Servers {
		if s.Enabled {
			wg.Add(1)
			go s.CreateAndConnect()
			infof("Connecting to %s", s.Network)
		} else {
			infof("Not connecting to %s because Enabled is false", s.Network)
		}
	}

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

	wg.Wait()
}

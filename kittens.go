package main

import (
	"os"
	"sync"
)

var (
	config *Config
	err    error
	wg     sync.WaitGroup
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

	wg.Wait()
}


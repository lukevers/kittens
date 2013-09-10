package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	// Name is a string that defines the default name of the bot
	// that will be connection to the IRC server.
	Name string

	// Server is a struct which contains information for the
	// servers that it connects to.
	Server struct {
		// Name is a string which defines the name of the
		// server that the bot is connecting to.
		Name string
		
		// Network is a string which defines the physical link
		// that the bot should try connection to.
		Network string
		
		// Port is an integer that defines the port that is
		// used to connect to the IRC server.
		Port int
		
		// Channels is a slice of strings that defines what
		// channels the bot should join upon connecting to the
		// server.
		Channels []string
	}
}

// ReadConfig reads the configuration file from JSON and returns it in
// the form of a *Config.
func ReadConfig(path string) (config *Config, err error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}

	config = &Config{}
	err = json.NewDecoder(file).Decode(config)

	return
}

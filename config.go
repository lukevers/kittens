package main

import (
	"encoding/json"
	"os"
)

var (
	config  *Config
)

type Config struct {
	// Debug is set to true if verbose messages are wanted, and
	// set to false if verbose messages are not wanted.
	Debug bool

	// The port that the webserver should lisen to incomming connections on
	// the default is 3000
	Port int

	// The interface that the webserver should listen to incomming
	// connections on. The default in our example config file is 0.0.0.0
	Interface string

	// Username is a string that contains the username for the user
	// to login and use the web interface.
	Username string

	// Password is a string that contains the password for the user
	// to login and use the web interface.
	Password string

	// Servers is a slice of Server structs. Kittens can connect
	// to multiple servers, and each server is defined in a new
	// Server struct.
	Servers []Server
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

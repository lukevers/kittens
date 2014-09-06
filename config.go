package main

import (
	"encoding/json"
	"os"
)

var (
	config *Config
)

type Config struct {
	// Debug is set to true if verbose messages are wanted, and
	// set to false if verbose messages are not wanted.
	Debug bool

	// DB is
	DB Database
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

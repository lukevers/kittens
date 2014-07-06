package main

import (
	irc "github.com/fluffle/goirc/client"
	"encoding/json"
	"os"
)

type Config struct {
	// Debug is set to true if verbose messages are wanted, and
	// set to false if verbose messages are not wanted.
	Debug bool

	// The port that the webserver should lisen to incomming connections on
	Port int

	// The interface that the webserver should listen to incomming
	// connections on. The default in our example config file is 0.0.0.0
	Interface string

	// Servers is a slice of Server structs. Kittens can connect
	// to multiple servers, and each server is defined in a new
	// Server struct.
	Servers []Server
}

type Server struct {

	// Conn is the connection that each bot is using to connect
	// to the server.
	Conn *irc.Conn

	// A unique ID will be given to each server when a goroutine
	// commences for the first time. This is used to identify
	// POST requests from our webinterface.
	ID uint16

	// Timestamp is a unix timestamp which will be set to time.Now
	// when the bot connects to the server.
	Timestamp int64

	// Nick is a string that defines the nick of the bot for this
	// specific server.
	Nick string

	// RealName is a string that defines the real name of the bot
	// for this specific server.
	RealName string

	// Host is a string that defines the host of the bot for this
	// specific server.
	Host string

	// ServerName is a string that defines the name of the server
	// that the bot is connecting to. (eg. freenode)
	ServerName string

	// Network is a string that defines the physical link that is
	// going to be used to connect to.
	Network string

	// Port is a number that defines the port that the bot uses
	// to connect to.
	Port int

	// SSL is set to true if the bot is connecting via SSL, and
	// set to false if the bot is not connecting via SSL.
	SSL bool

	// Password is a string that is only used if connecting to
	// the network requires a password.
	Password string

	// Enabled is set to true if the bot is currently enabled,
	// and set to false if it is not enabled.
	Enabled bool

	// Connected is set to true when the bot connects to the
	// server and set to false when it disconnects.
	Connected bool

	// Channels is a slice of strings that define what channels
	// the bot connects to.
	Channels []string
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

package main

import (
	irc "github.com/fluffle/goirc/client"
	"github.com/fluffle/goevent/event"
)

var (
	r event.EventRegistry
)

func main() {

	config, err := ReadConfig("example.config.json")
	if err != nil {
		panic(err)
	}
	
	conn := irc.Client(config.Nick, config.Host, config.Name, r)
	err = conn.Connect(config.Server.Network)
}
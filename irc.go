package main

import (
	"github.com/lukevers/goirc/client"
	"github.com/fluffle/goevent/event"
)

var (
	r event.EventRegistry
)

// CreateBot is a func that creates the physical bot that will connect
// to the server.
func CreateBot(config *Config) *client.Conn {
	l.Infof("Creating %s the bot", config.Nick)

	conf := client.NewConfig(config.Nick, []string{config.Host, config.Name}...)
	
	// Check for SSL
	if config.Server.SSL {
		conf.SSL = true
	}

	// Get the server/port we are going to connect to
	if config.Server.Port != 6667 {
		conf.Server = config.Server.Network + ":" + string(config.Server.Port)
	} else {
		conf.Server = config.Server.Network
	}

	// Create the new client with the configurations we added to
	// *client.Config from our *Config.
	bot := client.Client(conf)
	
	// Enable state tracking
	bot.EnableStateTracking()
	
	l.Infof("Created %s the bot", config.Nick)
	return bot
}

// Connect is a func that connects to the server and stays connected
// until being disconnected.
func Connect(bot *client.Conn, config *Config) {

	// Join channels on connect
	bot.HandleFunc("connected",
		func (conn *client.Conn, line *client.Line) {
			JoinChannels(bot, config)
		})
	
	// Handler for disconnect events
	quit := make(chan bool)
	bot.HandleFunc("disconnected",
		func(conn *client.Conn, line *client.Line) { 
			quit <- true 
		})
	
	// Connect to server
	l.Infof("Connecting to %s", config.Server.Name)
	
	if err := bot.Connect(); err != nil {
		l.Emergf("Connection error: %s", err)
	}
	
	// Wait for quit
	<-quit
}

// JoinChannels is a func that is called before connecting to the
// server so it knows what channels to connect to.
func JoinChannels(bot *client.Conn, config *Config) {
	for i := range config.Server.Channels {
		l.Infof("Joining channel %s", config.Server.Channels[i])
		bot.Join(config.Server.Channels[i])
	}
}
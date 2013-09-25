package main

import (
	"github.com/fluffle/goirc/client"
	"github.com/fluffle/goevent/event"
)

var (
	r event.EventRegistry
)

// CreateBot is a function that creates the physical bot that will
// connect to the server.
func CreateBot(config *Config) *client.Conn {
	l.Infof("Creating %s the bot", config.Nick)
	r = event.NewRegistry()
	bot := client.Client(config.Nick, config.Host, config.Name, r)
	if config.Server.SSL {
		bot.SSL = true
	}
	bot.EnableStateTracking()
	l.Infof("Created %s the bot", config.Nick)
	return bot
}

// Connect is a function that connects to the server and joins all of
// the channels that are set in the configuration file.
func Connect(bot *client.Conn, config *Config) {
	l.Infof("Preparing to connect to %s", config.Server.Network)
	// Join channels
	bot.AddHandler(client.CONNECTED, 
		func (conn *client.Conn, line *client.Line) { 
			JoinChannels(bot, config) 
		})

	quit := make(chan bool)

	bot.AddHandler(client.DISCONNECTED, 
		func(conn *client.Conn, line *client.Line) { 
			quit <- true 
		})
	
	l.Infof("Connecting to %s", config.Server.Network)
	// Connect to server
	if err := bot.Connect(config.Server.Network); err != nil {
		l.Fatalf("Error: %s", err)
	}

	<-quit
}

// JoinChannels is a function that is called before connecting to the
// server so it knows what channels to connect to.
func JoinChannels(bot *client.Conn, config *Config) {
	for i := range config.Server.Channels {
		l.Infof("Joining channel %s", config.Server.Channels[i])
		bot.Join(config.Server.Channels[i])
	}
}

func AddHandler(bot *client.Conn) *client.Conn {
	
	
		
	return bot
}
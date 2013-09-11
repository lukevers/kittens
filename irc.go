package main

import (
	"github.com/fluffle/goirc/client"
	"github.com/fluffle/goevent/event"
)

var (
	r event.EventRegistry
)

func CreateBot(config *Config) *client.Conn {
	r = event.NewRegistry()
	bot := client.Client(config.Nick, config.Host, config.Name, r)
	if config.Server.SSL {
		bot.SSL = true
	}
	return bot
}

func Connect(bot *client.Conn, config *Config) {

	bot.AddHandler(client.CONNECTED, func(conn *client.Conn, line *client.Line) {
		for i := range config.Server.Channels {
			bot.Join(config.Server.Channels[i])
		}
	})

	quit := make(chan bool)

	bot.AddHandler(client.DISCONNECTED, func(conn *client.Conn, line *client.Line) { quit <- true })

	// Connect to server
	if err := bot.Connect(config.Server.Network); err != nil {
		panic(err)
	}

	<-quit
}

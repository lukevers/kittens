package main

import (
	"github.com/fluffle/goevent/event"
	irc "github.com/fluffle/goirc/client"
	"strconv"
	"time"
)

func (s Server) CreateAndConnect() {
	cli = append(cli, &s)
	verbf("Creating bot from server struct: %s", s)

	r := event.NewRegistry()
	conn := irc.Client(s.Nick, s.Host, s.RealName, r)

	// Set our SSL setting
	conn.SSL = s.SSL

	verbf("Finished creating bot for server %s", s.Network)
	verbf("Beginning to connect to %s", s.Network)

	// Register connect handler
	conn.AddHandler(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			s.Connected = true
			infof("Connected to %s", s.Network)
			s.JoinChannels(conn)
		})

	quit := make(chan bool)

	// Register disconnect handler
	conn.AddHandler(irc.DISCONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			s.Connected = false
			infof("Disconnected from %s", s.Network)
			quit <- true
		})

	// Register plugin handlers
	s.AddPlugins(conn)

	// Now we connect
	if err := conn.Connect(s.Network+":"+strconv.Itoa(s.Port), s.Password); err != nil {
		warnf("Error connecting: %s", err)
		info("Retrying in 30 seconds")
		time.Sleep(30 * time.Second)
		go s.CreateAndConnect()
	}

	// Wait for disconnect
	<-quit
}

// JoinChannels is a func that is called when a bot connects
// to a server. The func loops over the channels that are in
// the slice of channels in our Server struct.
func (s Server) JoinChannels(conn *irc.Conn) {
	for i := range s.Channels {
		verbf("Joining channel: %s", s.Channels[i])
		conn.Join(s.Channels[i])
	}
}

// 
func (s Server) AddPlugins (conn *irc.Conn) {
	
}

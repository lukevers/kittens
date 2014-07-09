package main

import (
	"github.com/fluffle/goevent/event"
	irc "github.com/fluffle/goirc/client"
	"reflect"
	"strconv"
	"time"
)

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

	// Channels is a slice of Channel Structs that define what channels
	// the bot connects to.
	Channels []*Channel
}

func (s Server) CreateAndConnect(new bool) {
	if !new {
		for i, v := range clients {
			if reflect.DeepEqual(&s, v) {
				clients[i] = nil
				clients[i] = &s
			}
		}
	} else {
		clients = append(clients, &s)
	}

	verbf("Creating bot from server struct: %s", s)

	r := event.NewRegistry()
	s.Conn = irc.Client(s.Nick, s.Host, s.RealName, r)

	// Set our SSL setting
	s.Conn.SSL = s.SSL

	// Set our PING Frequency to a lower time than default
	s.Conn.PingFreq = (30 * time.Second)

	verbf("Finished creating bot for server %s", s.ServerName)
	verbf("Beginning to connect to %s", s.Network)

	// Register connect handler
	s.Conn.AddHandler(irc.CONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			s.Timestamp = time.Now().Unix()
			s.Connected = true
			infof("Connected to %s", s.Network)
			s.JoinChannels()
		})

	quit := make(chan bool)

	// Register disconnect handler
	s.Conn.AddHandler(irc.DISCONNECTED,
		func(conn *irc.Conn, line *irc.Line) {
			s.Connected = false
			infof("Disconnected from %s", s.Network)
			infof("Reconnecting to %s", s.Network)
			go s.CreateAndConnect(false)
			quit <- true
			close(quit)
		})

	s.Conn.AddHandler("PRIVMSG",
		func(conn *irc.Conn, line *irc.Line) {
			// Show output of line currently
			s.Logging(line)
		})

	// Now we connect
	if s.Enabled {
		if err := s.Conn.Connect(s.Network+":"+strconv.Itoa(s.Port), s.Password); err != nil {
			warnf("Error connecting: %s", err)
			info("Retrying in 30 seconds")
			time.Sleep(30 * time.Second)
			go s.CreateAndConnect(false)
			quit <- true
			close(quit)
		}
	} else {
		infof("Not connecting to %s because enabled is false", s.ServerName)
	}

	// Wait for disconnect
	<-quit
}

// Join Channels is a func that is called when a bot connects
// to a server. The func loops over the channels that are in
// the slice of channels in our Server struct.
func (s *Server) JoinChannels() {
	for i := range s.Channels {
		verbf("Joining channel: %s", s.Channels[i].Name)
		s.Conn.Join(s.Channels[i].Name)
	}
}

// Join New Channel is a func that is called when the bot is
// joining one specific channel for the first time.
func (s *Server) JoinNewChannel(channel string) {
	s.Channels = append(s.Channels, &Channel{
		Name: channel,
	})

	verbf("Joining channel: %s", channel)
	s.Conn.Join(channel)
}

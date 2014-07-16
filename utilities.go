package main

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Get Server From Request takes a *http.Request in and returns
// the server ID that is being looked at. It returns an error if
// it can't find it.
func GetServerFromRequest(req *http.Request) (*Server, error) {
	// Figure out what {id} is in "/server/{id}"
	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 16)
	if err != nil {
		warnf("Error converting server id: %s", err)
	}

	// Get our server from our slice of servers
	for _, s := range clients {
		if s.ID == uint16(id) {
			return s, nil
		}
	}

	return nil, errors.New("Could not find server")
}

// Get Channel From Request takes a *Server and *http.Request
// and returns the *Channel that is being looked at. It returns
// and error if it can't find it.
func GetChannelFromRequest(s *Server, req *http.Request) (*Channel, error) {
	// Figure out what {channel} is
	channel := mux.Vars(req)["channel"]

	// Get our channel from our slice of channels
	for _, c := range s.Channels {
		if channel == c.Name {
			return c, nil
		}
	}

	return nil, errors.New("Could not find channel")
}

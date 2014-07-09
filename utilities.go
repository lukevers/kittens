package main

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Get Server From Request takes a *http.Request in and returns
// the server ID that is being looked at. It returns an error if
// it can't find it somehow.
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

package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Handle POST requests to "/server/{id}" which are server update
// requests. From here we also want to update the live bot.
func UpdateServer(w http.ResponseWriter, req *http.Request) {

	// Figure out what {id} is in "/server/{id}"
	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 16)
	if err != nil {
		warnf("Error converting server id: %s", err)
	}

	// Get our server from our slice of servers
	var server *Server
	for _, s := range clients {
		if s.ID == uint16(id) {
			server = s
		}
	}

	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	// Check if Nick has been changed, and update it if it has
	if req.Form["nick"][0] != server.Nick {
		verbf("Changing nick to %s", req.Form["nick"][0])
		server.Nick = req.Form["nick"][0]
		server.Conn.Nick(server.Nick)
	}

	// Redirect back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), 303)
}

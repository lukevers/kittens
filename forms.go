package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Helper func so less code is duplicated here.
func GetDefault(req *http.Request) (server *Server, id uint16) {
	// Figure out what {id} is in "/server/{id}"
	id_, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 16)
	if err != nil {
		warnf("Error converting server id: %s", err)
	} else {
		id = uint16(id_)
	}

	// Get our server from our slice of servers
	for _, s := range clients {
		if s.ID == uint16(id) {
			server = s
		}
	}

	return
}

// Handle POST requests to "/server/{id}" which are server update
// requests. From here we also want to update the live bot.
func UpdateServer(w http.ResponseWriter, req *http.Request) {

	server, _ := GetDefault(req)

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

	// Check if Real Name has been changed, and update it if it has
	if req.Form["realname"][0] != server.RealName {
		verbf("Changing real name to %s", req.Form["realname"][0])
		server.RealName = req.Form["realname"][0]
		server.Conn.Raw("SETNAME " + server.RealName)
	}

	// Check if Host has been changed, and update it if it has
	if req.Form["host"][0] != server.Host {
		verbf("Changing host to %s", req.Form["host"][0])
		server.Host = req.Form["host"][0]
	}

	// Check if Server Name has been changed, and update it if it has
	if req.Form["servername"][0] != server.ServerName {
		verbf("Changing server name to %s", req.Form["servername"][0])
		server.ServerName = req.Form["servername"][0]
	}

	// Check if Network has been changed, and update it if it has
	if req.Form["network"][0] != server.Network {
		verbf("Changing network to %s", req.Form["network"][0])
		server.Network = req.Form["network"][0]
	}

	// Check if Port has been changed, and update it if it has
	p, err := strconv.Atoi(req.Form["port"][0])
	if err != nil {
		warnf("Error converting Port from form to int: %s", err)
	}
	if p != server.Port {
		verbf("Changing port to %s", req.Form["port"][0])
		server.Port = p
	}

	// Check if Password has been changed, and update it if it has
	if req.Form["password"][0] != server.Password {
		verbf("Changing password to %s", req.Form["password"][0])
		server.Password = req.Form["password"][0]
	}

	// Redirect (303) back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), http.StatusSeeOther)
}

// Handle POST requests to "/server/{id}/channel/join" which are
// server join channel requests. From here we also want to update
// the live bot.
func JoinChannel(w http.ResponseWriter, req *http.Request) {

	server, _ := GetDefault(req)

	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	ch := req.Form["channel"][0]

	copied := false
	for _, v := range server.Channels {
		if ch == v {
			copied = true
		}
	}

	if !copied {
		server.Channels = append(server.Channels, ch)
		verbf("Joining channel %s", ch)
		server.Conn.Join(ch)
	}

	// Redirect (303) back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), http.StatusSeeOther)
}

// Handle POST requests to "/server/{id}/channel/part" which are
// server part channel requests. From here we also want to update
// the live bot.
func PartChannel(w http.ResponseWriter, req *http.Request) {

	server, _ := GetDefault(req)

	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	ch := req.Form["channel"][0]
	for i, v := range server.Channels {
		if ch == v {
			server.Channels = append(server.Channels[:i], server.Channels[i+1:]...)
		}
	}

	verbf("Parting channel %s", ch)
	server.Conn.Part(ch)

	// Redirect (303) back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), http.StatusSeeOther)
}

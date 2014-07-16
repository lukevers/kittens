package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Handle "/" web
func HandleRoot(w http.ResponseWriter, req *http.Request) {
	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	templates.ExecuteTemplate(w, "index", clients)
}

// Handle "/server/{id}" web
func HandleServer(w http.ResponseWriter, req *http.Request) {
	server, err := GetServerFromRequest(req)
	if err != nil {
		warnf("Error parsing server request: %s", err)
	}

	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	templates.ExecuteTemplate(w, "server", server)
}

// Handle "/server/{id}/channel/{channel}" web
func HandleChannel(w http.ResponseWriter, req *http.Request) {
	server, err := GetServerFromRequest(req)
	if err != nil {
		warnf("Error parsing server request: %s", err)
	}

	channel, err := GetChannelFromRequest(server, req)
	if err != nil {
		warnf("Error parsing channel request: %s", err)
	}

	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	// Template data. We want to pass both the *Server for server
	// information, and the *Channel so we don't have to loop
	// through the slice of *Channels each time we want to access
	// our *Channel. We're passing an anonymous struct to do this.
	data := struct {
		Server *Server
		Channel *Channel
	}{
		server,
		channel,
	}

	templates.ExecuteTemplate(w, "channel", data)
}

// Handle "/server/{id}/channel/" web
func HandleChannelRedirect(w http.ResponseWriter, req *http.Request) {
	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	templates.ExecuteTemplate(w, "redirect", nil)
}

// Handle POST requests to "/server/{id}/channel/join" which are
// server join channel requests. From here we also want to update
// the live bot.
func HandleJoinChannel(w http.ResponseWriter, req *http.Request) {
	server, err := GetServerFromRequest(req)
	if err != nil {
		warnf("Error parsing server request: %s", err)
	}

	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	ch := req.Form["channel"][0]

	// Check to see if we're already in the channel we are trying
	// to currently join.
	copied := false
	for _, v := range server.Channels {
		if ch == v.Name {
			copied = true
			break
		}
	}

	// If we're not trying to join a channel we're already in
	// let's join that channel.
	if !copied {
		server.JoinNewChannel(ch)
	}

	// Redirect (303) back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), http.StatusSeeOther)
}

// Handle POST requests to "/server/{id}/channel/part" which are
// server part channel requests. From here we also want to update
// the live bot.
func HandlePartChannel(w http.ResponseWriter, req *http.Request) {
	server, err := GetServerFromRequest(req)
	if err != nil {
		warnf("Error parsing server request: %s", err)
	}

	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	ch := req.Form["channel"][0]

	for i, v := range server.Channels {
		if ch == v.Name {
			server.Channels = append(server.Channels[:i], server.Channels[i+1:]...)
		}
	}

	verbf("Parting channel %s", ch)
	server.Conn.Part(ch)

	// Redirect (303) back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), http.StatusSeeOther)
}

// Handle POST requests to "/server/{id}" which are server update
// requests. From here we also want to update the live bot.
func HandleUpdateServer(w http.ResponseWriter, req *http.Request) {
	server, err := GetServerFromRequest(req)
	if err != nil {
		warnf("Error parsing server request: %s", err)
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

func HandleEnableServer(w http.ResponseWriter, req *http.Request) {
	server, err := GetServerFromRequest(req)
	if err != nil {
		warnf("Error parsing server request: %s", err)
	}

	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	enabled, err := strconv.ParseBool(req.Form["enabled"][0])
	if err != nil {
		warnf("Error parsing enabled from string to bool: %s", err)
	}

	if enabled {
		// Enable and connect
		server.Enabled = true
		go server.CreateAndConnect(false)
	} else {
		// Disable and disconnect
		server.Enabled = false
		server.Conn.Quit()
	}

	// Redirect (303) back to "/server/{id}" when we're done here
	http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.ID)), http.StatusSeeOther)
}

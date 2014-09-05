package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Handle "/" web
func HandleRoot(w http.ResponseWriter, req *http.Request) {
	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions(req)).ParseGlob("app/views/*"))
	}

	if IsLoggedIn(req) {
		templates.Funcs(AddTemplateFunctions(req)).ExecuteTemplate(w, "index", WhoAmI(req))
	} else {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	}
}

// Handle "/logout" web
func HandleLogout(w http.ResponseWriter, req *http.Request) {
	// Remove cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	// Redirect to "/"
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// Handle "/login" web
func HandleLogin(w http.ResponseWriter, req *http.Request) {
	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions(req)).ParseGlob("app/views/*"))
	}

	if IsLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else {
		templates.Funcs(AddTemplateFunctions(req)).ExecuteTemplate(w, "login", nil)
	}
}

// Handle POSTS to "/login" web.
func HandleLoginForm(w http.ResponseWriter, req *http.Request) {
	// Parse our form so we can get values from req.Form
	err = req.ParseForm()
	if err != nil {
		warnf("Error parsing form: %s", err)
	}

	// Get username/password from input
	username := req.Form["username"][0]
	password := req.Form["password"][0]

	// Query database for user
	var user User
	db.Table("users").Where("username = ?", username).First(&user)

	// Check if usernames match up
	if user.Username == username {
		// Check if passwords match up
		if PasswordMatchesHash(password, user.Password) {
			// Create new session
			session, err := store.New(req, "user")
			session.Values["username"] = username
			if err != nil {
				warnf("Error creating new session: %s", err)
			}

			// Save session and redirect
			session.Save(req, w)
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
	}

	// If you have gotten this far then you have not been
	// authenticated. Sorry.
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

// Handle "/server/{id}" web
func HandleServer(w http.ResponseWriter, req *http.Request) {
	// Check if logged in
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
			// Refresh the templates
			if config.Debug {
				templates = RefreshTemplates(req)
			}

			// Execute template
			templates.Funcs(AddTemplateFunctions(req)).ExecuteTemplate(w, "server", server)
		}
	}
}

// Handle "/server/{id}/channel/{channel}" web
func HandleChannel(w http.ResponseWriter, req *http.Request) {
	// Check if logged in
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
			// Get channel from request
			channel, err := GetChannelFromRequest(server, req)
			if err != nil {
				warnf("Error parsing channel request: %s", err)
			}

			// Refresh the templates
			if config.Debug {
				templates = RefreshTemplates(req)
			}

			// Template data. We want to pass both the *Server for server
			// information, and the *Channel so we don't have to loop
			// through the slice of *Channels each time we want to access
			// our *Channel. We're passing an anonymous struct to do this.
			data := struct {
				Server  *Server
				Channel *Channel
			}{
				server,
				channel,
			}

			// Execute template
			templates.Funcs(AddTemplateFunctions(req)).ExecuteTemplate(w, "channel", data)
		}
	}
}

// Handle "/server/{id}/channel/" web
func HandleChannelRedirect(w http.ResponseWriter, req *http.Request) {
	// Check if logged in
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
			// Refresh the templates
			if config.Debug {
				templates = RefreshTemplates(req)
			}

			// Execute Template
			templates.Funcs(AddTemplateFunctions(req)).ExecuteTemplate(w, "redirect", nil)
		}
	}
}

// Handle POST requests to "/server/{id}/channel/join" which are
// server join channel requests. From here we also want to update
// the live bot.
func HandleJoinChannel(w http.ResponseWriter, req *http.Request) {
	// Check if logged in
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
			// Parse our form so we can get values from req.Form
			err = req.ParseForm()
			if err != nil {
				warnf("Error parsing form: %s", err)
			}

			// Get our channel name
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

			// Redirect back to "/server/{id}" when we're done here
			http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.Id)), http.StatusSeeOther)
		}
	}
}

// Handle POST requests to "/server/{id}/channel/part" which are
// server part channel requests. From here we also want to update
// the live bot.
func HandlePartChannel(w http.ResponseWriter, req *http.Request) {
	// Check if logged in
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
			// Parse our form so we can get values from req.Form
			err = req.ParseForm()
			if err != nil {
				warnf("Error parsing form: %s", err)
			}

			// Get our channel name
			ch := req.Form["channel"][0]

			// Loop through channels and when we find it, remove it.
			for i, v := range server.Channels {
				if ch == v.Name {
					// Delete from database
					db.Unscoped().Table("channels").Where("id = ?", v.Id).Delete(&Channel{})

					// Remove from structs
					server.Channels = append(server.Channels[:i], server.Channels[i+1:]...)
				}
			}

			// Part channel
			verbf("Parting channel %s", ch)
			server.Conn.Part(ch)

			// Redirect back to "/server/{id}" when we're done here
			http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.Id)), http.StatusSeeOther)
		}
	}
}

// Handle POST requests to "/server/{id}" which are server update
// requests. From here we also want to update the live bot.
func HandleUpdateServer(w http.ResponseWriter, req *http.Request) {
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
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

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("nick", server.Nick)
			}

			// Check if Real Name has been changed, and update it if it has
			if req.Form["realname"][0] != server.RealName {
				verbf("Changing real name to %s", req.Form["realname"][0])
				server.RealName = req.Form["realname"][0]
				server.Conn.Raw("SETNAME " + server.RealName)

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("real_name", server.RealName)
			}

			// Check if Host has been changed, and update it if it has
			if req.Form["host"][0] != server.Host {
				verbf("Changing host to %s", req.Form["host"][0])
				server.Host = req.Form["host"][0]

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("host", server.Host)
			}

			// Check if Server Name has been changed, and update it if it has
			if req.Form["servername"][0] != server.ServerName {
				verbf("Changing server name to %s", req.Form["servername"][0])
				server.ServerName = req.Form["servername"][0]

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("server_name", server.ServerName)
			}

			// Check if Network has been changed, and update it if it has
			if req.Form["network"][0] != server.Network {
				verbf("Changing network to %s", req.Form["network"][0])
				server.Network = req.Form["network"][0]

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("network", server.Network)
			}

			// Check if Port has been changed, and update it if it has
			p, err := strconv.Atoi(req.Form["port"][0])
			if err != nil {
				warnf("Error converting Port from form to int: %s", err)
			}
			if p != server.Port {
				verbf("Changing port to %s", req.Form["port"][0])
				server.Port = p

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("port", server.Port)
			}

			// Check if Password has been changed, and update it if it has
			if req.Form["password"][0] != server.Password {
				verbf("Changing password to %s", req.Form["password"][0])
				server.Password = req.Form["password"][0]

				// Update in database
				db.Table("servers").Where("id = ?", server.Id).Update("password", server.Password)
			}

			// Redirect back to "/server/{id}" when we're done here
			http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.Id)), http.StatusSeeOther)
		}
	}
}

// Handles POST requests for "/sever/{id}/enable" which either enables
// or disables a server depending on if the server is currently enabled
// or disabled, and if the bool we are given is true or false.
func HandleEnableServer(w http.ResponseWriter, req *http.Request) {
	// Check if logged in
	if !IsLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	} else {
		// Get server from request
		server, err := GetServerFromRequest(req)
		if err != nil {
			warnf("Error parsing server request: %s", err)
		}

		// Check if user owns the server
		if !WhoAmI(req).OwnsServer(server) {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		} else {
			// Parse our form so we can get values from req.Form
			err = req.ParseForm()
			if err != nil {
				warnf("Error parsing form: %s", err)
			}

			// Get form data "enabled" and convert it to a bool
			enabled, err := strconv.ParseBool(req.Form["enabled"][0])
			if err != nil {
				warnf("Error parsing enabled from string to bool: %s", err)
			}

			// Get server from database to prepare to update enabled
			var s Server
			db.First(&s, server.Id)

			// Check if enabled or not
			if enabled {
				// Enable and connect
				server.Enabled = true
				go server.CreateAndConnect(false)
			} else {
				// Disable and disconnect
				server.Enabled = false
				server.Conn.Quit()
			}

			// Update database
			s.Enabled = server.Enabled
			db.Save(&s)

			// Redirect back to "/server/{id}" when we're done here
			http.Redirect(w, req, "/server/"+strconv.Itoa(int(server.Id)), http.StatusSeeOther)
		}
	}
}

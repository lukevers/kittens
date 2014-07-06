package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

// Template func that counts connected servers
func ConnectedServers() string {
	i := 0
	for _, s := range clients {
		if s.Connected {
			i++
		}
	}

	if i == 1 {
		return "1 Connected Server"
	} else {
		return strconv.Itoa(i) + " Connected Servers"
	}
}

// Template func that counts enabled servers
func EnabledServers() string {
	i := 0
	for _, s := range clients {
		if s.Enabled {
			i++
		}
	}

	if i == 1 {
		return "1 Enabled Server"
	} else {
		return strconv.Itoa(i) + " Enabled Servers"
	}
}

// Template func that counts disabled servers
func DisabledServers() string {
	i := 0
	for _, s := range clients {
		if !s.Enabled {
			i++
		}
	}

	if i == 1 {
		return "1 Disabled Server"
	} else {
		return strconv.Itoa(i) + " Disabled Servers"
	}
}

// Template func that counts total servers
func TotalServers() string {
	if len(clients) > 1 {
		return strconv.Itoa(len(clients)) + " Total Servers"
	} else {
		return "1 Total Server"
	}
}

// Add func to templates
func AddTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"EnabledServers":   EnabledServers,
		"TotalServers":     TotalServers,
		"ConnectedServers": ConnectedServers,
		"DisabledServers":  DisabledServers,
	}
}

// Handle "/" web
func HandleRoot(w http.ResponseWriter, req *http.Request) {
	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	templates.ExecuteTemplate(w, "index", clients)
}

// Handle "/server/{id}" web
func HandleServer(w http.ResponseWriter, req *http.Request) {

	id, err := strconv.ParseUint(mux.Vars(req)["id"], 10, 16)
	if err != nil {
		warnf("Error converting server id: %s", err)
	}

	var server *Server
	for _, s := range clients {
		if s.ID == uint16(id) {
			server = s
		}
	}

	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	templates.ExecuteTemplate(w, "server", server)
}

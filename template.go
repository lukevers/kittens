package main

import (
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

// Template func that counts total servers
func TotalServers() string {
	if (len(clients) > 1) {
		return strconv.Itoa(len(clients)) + " Total Servers"
	} else {
		return "1 Total Server"
	}
}

// Add func to templates
func AddTemplateFunctions() (template.FuncMap) {
	return template.FuncMap{
		"EnabledServers": EnabledServers,
		"TotalServers": TotalServers,
		"ConnectedServers": ConnectedServers,
	}
}

// Handle "/" web
func HandleRoot(w http.ResponseWriter, req *http.Request) {
	
	if config.Debug {
		templates = template.Must(template.New("").Funcs(AddTemplateFunctions()).ParseGlob("app/views/*"))
	}

	templates.ExecuteTemplate(w, "index", cli)
}

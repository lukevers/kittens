package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Template func that counts connected servers
func ConnectedServers(req *http.Request) string {
	user := WhoAmI(req)

	i := 0
	for _, s := range user.Servers {
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
func EnabledServers(req *http.Request) string {
	user := WhoAmI(req)

	i := 0
	for _, s := range user.Servers {
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
func DisabledServers(req *http.Request) string {
	user := WhoAmI(req)

	i := 0
	for _, s := range user.Servers {
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
func TotalServers(req *http.Request) string {
	user := WhoAmI(req)

	if len(user.Servers) != 1 {
		return strconv.Itoa(len(user.Servers)) + " Total Servers"
	} else {
		return "1 Total Server"
	}
}

// Add func to templates
func AddTemplateFunctions(req *http.Request) template.FuncMap {
	return template.FuncMap{
		"EnabledServers":   func() string { return EnabledServers(req) },
		"TotalServers":     func() string { return TotalServers(req) },
		"ConnectedServers": func() string { return ConnectedServers(req) },
		"DisabledServers":  func() string { return DisabledServers(req) },
	}
}

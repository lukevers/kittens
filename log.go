package main

import (
	irc "github.com/fluffle/goirc/client"
	"github.com/mgutz/ansi"
	"log"
	"os"
)

var (
	d = ansi.ColorCode("white+b:magenta")
	i = ansi.ColorCode("white+b:blue")
	w = ansi.ColorCode("white+b:red")
	v = ansi.ColorCode("white+b:black")
	r = ansi.ColorCode("reset")
)

var (
	INFO = log.New(os.Stdout, d+"[kittens]"+r+" "+i+"INFO"+r+" ", 0)
	WARN = log.New(os.Stdout, d+"[kittens]"+r+" "+w+"WARN"+r+" ", 0)
	VERB = log.New(os.Stdout, d+"[kittens]"+r+" "+v+"VERB"+r+" ", 0)
)

// Print info statements
func info(i interface{}) {
	INFO.Print(i)
}

// Printf info statements
func infof(s string, i interface{}) {
	INFO.Printf(s, i)
}

// Print warning statements
func warn(i interface{}) {
	WARN.Print(i)
}

// Printf warning statements
func warnf(s string, i interface{}) {
	WARN.Printf(s, i)
}

// Print verbose statements if debug is true
func verb(i interface{}) {
	if config.Debug {
		VERB.Print(i)
	}
}

// Printf verbose statements if debug is true
func verbf(s string, i interface{}) {
	if config.Debug {
		VERB.Printf(s, i)
	}
}

// The logging func logs chat information.
//
// For reference:
// *irc.Line
//
// type Line struct {
//     Nick, Ident, Host, Src string
//     Cmd, Raw string
//     Args []string
//     Time time.Time
// }
//
func (s Server) Logging(line *irc.Line) {
	// Our standard permissions
	var perms os.FileMode = 0760

	// Get some info for easier access
	channel := line.Args[0]

	// If we don't have our "logs" folder, create it
	if _, err := os.Stat("./logs"); err != nil {
		if os.IsNotExist(err) {
			verb("Creating directory \"logs\"")
			os.Mkdir("./logs", perms)
		}
	}

	// If we don't have our "logs/{server}" folder, create it
	if _, err := os.Stat("./logs/" + s.ServerName); err != nil {
		if os.IsNotExist(err) {
			verb("Creating directory \"logs/" + s.ServerName + "\"")
			os.Mkdir("./logs/"+s.ServerName, perms)
		}
	}

	// If we don't have our "logs/{server}/{channel}" folder,
	// create it.
	if _, err := os.Stat("./logs/" + s.ServerName + "/" + channel); err != nil {
		if os.IsNotExist(err) {
			verb("Creating directory \"logs/" + s.ServerName + "/" + channel + "\"")
			os.Mkdir("./logs/"+s.ServerName+"/"+channel, perms)
		}
	}
}

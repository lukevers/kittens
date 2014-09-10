package main

import (
	"flag"
)

var (
	// General flags
	debugFlag = flag.Bool("debug", false, "Use during development, not production")

	// Webserver flags
	portFlag      = flag.Int("port", 3000, "Port for webserver to bind to")
	interfaceFlag = flag.String("interface", "0.0.0.0", "Interface for webserver to bind to")

	// Database flags
	driverFlag   = flag.String("driver", "sqlite", "Database driver")
	databaseFlag = flag.String("database", "kittens.db", "Database string")
)

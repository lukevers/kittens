package main

import (
	"flag"
)

var (
	// Path to configuration file flags
	configPathFlag = flag.String("config", "config.json", "Path to configuration file")

	// General flags
	debugFlag = flag.Bool("debug", false, "Use during development, not production")

	// Webserver flags
	portFlag      = flag.Int("port", 3000, "Port for webserver to bind to")
	interfaceFlag = flag.String("interface", "0.0.0.0", "Interface for webserver to bind to")
	noAuthFlag    = flag.Bool("no-auth", false, "Turn off login system. For development purposes mainly")
)

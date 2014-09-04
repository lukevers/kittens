package main

import (
	"flag"
)

var (
	// Path to configuration file
	configPathFlag = flag.String("config", "config.json", "Path to configuration file")

	// Configuration flags
	debugFlag     = flag.Bool("debug", false, "Use during development, not production")
	portFlag      = flag.Int("port", 3000, "Port for webserver to bind to")
	interfaceFlag = flag.String("interface", "0.0.0.0", "Interface for webserver to bind to")
)

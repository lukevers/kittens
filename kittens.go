package main

import (
	"github.com/inhies/go-log"
	"fmt"
	"os"
)

var (
	l *log.Logger
	LogLevel = log.LogLevel(log.INFO)
	LogFlags = log.Ldate | log.Ltime
	LogFile  = os.Stdout
)

func main() {
	
	// Start the logger
	l, err := log.NewLevel(LogLevel, true, LogFile, "", LogFlags)
	if err != nil {
		fmt.Printf("Could not start logger: %s", err)
		os.Exit(1)
	}

	// Load the configuration file
	config, err := ReadConfig("example.config.json")
	if err != nil {
		l.Fatalf("Could not read configuration file: %s", err)
	}

	// Create the bot
	bot := CreateBot(config)
		
	// Add listeners
	// todo
	
	// Connect to server
	Connect(bot, config)
}
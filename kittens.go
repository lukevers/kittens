package main

import (
	"fmt"
)

func main() {

	config, err := ReadConfig("example.config.json")

	if err != nil {
		fmt.Printf("Can not create config: %s\n",err)
	}

	// Create the bot
	bot := CreateBot(config)
		
	// Add listeners
	// todo
	
	// Connect to server
	Connect(bot, config)
}
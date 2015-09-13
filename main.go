package main

import (
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	InitDatabase()
	InitPluginAPI()
	InitRouter()
}

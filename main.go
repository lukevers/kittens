/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/davecgh/go-spew/spew"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	InitDatabase()
	InitEnabledBots()
	InitRouter()
}

func dump(i interface{}) {
	spew.Dump(i)
}

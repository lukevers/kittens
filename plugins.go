/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/thoj/go-ircevent"
	"github.com/yuin/gopher-lua"
)

var bot *irc.Connection
var L *lua.LState

func createbot() {
	bot = irc.IRC("bot", "botuser")
	bot.Connect("localhost:6667")
	bot.Join("#test")

	if err := L.DoFile("plugins/say.lua"); err != nil {
		panic(err)
	}

	if err := L.DoFile("plugins/echo.lua"); err != nil {
		panic(err)
	}
}

func InitPluginAPI() {
	// And now we can work on the plugin stuff
	L = lua.NewState()
	defer L.Close()

	L.SetGlobal("say", L.NewFunction(Say))
	L.SetGlobal("on", L.NewFunction(On))

	// For now let's create an example bot
	createbot()
}

func On(L *lua.LState) int {
	event := L.ToString(1)
	//	cback := L.ToString(2)

	bot.AddCallback(event, func(event *irc.Event) {
		dump(event)

		/*
			        L.DoString(fmt.Sprintf(`%s("%s", "%s")`,
						cback,
						event.Arguments[0],
						event.Message()))*/
	})

	return 1
}

func Say(L *lua.LState) int {
	channel := L.ToString(1)
	message := L.ToString(2)

	bot.Privmsg(channel, message)

	return 1
}

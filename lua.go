/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/yuin/gopher-lua"
	"github.com/thoj/go-ircevent"
	"fmt"
)

type Lua struct {
	Lua *lua.LState
	Bot *Bot
}

func NewLuaState(bot *Bot) *Lua {
	L := &Lua{
		Lua: lua.NewState(),
		Bot: bot,
	}

	defer L.Lua.Close()
	return L.SetPluginAPI()
}

func (L *Lua) SetPluginAPI() *Lua {
	L.Lua.SetGlobal("say", L.Lua.NewFunction(L.say))
	L.Lua.SetGlobal("on", L.Lua.NewFunction(L.on))
	return L
}

func (L *Lua) say(state *lua.LState) int {
	channel := state.ToString(1)
	message := state.ToString(2)
	L.Bot.irc.Privmsg(channel, message)

	return 1
}

func (L *Lua) on(state *lua.LState) int {
	event := state.ToString(1)
	cback := state.ToString(2)

	L.Bot.irc.AddCallback(event, func(event *irc.Event) {
		state.DoString(fmt.Sprintf(`%s("%s", "%s")`,
			cback,
			event.Arguments[0],
			event.Message()))
	})

	return 1
}

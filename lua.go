/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	LUA "github.com/yuin/gopher-lua"
	"github.com/thoj/go-ircevent"
	"fmt"
)

type Lua struct {
	L *LUA.LState
	B *Bot
}

func NewLuaState(bot *Bot) *Lua {
	lua := &Lua{
		L: LUA.NewState(),
		B: bot,
	}

	defer lua.L.Close()
	return lua.SetPluginAPI()
}

func (lua *Lua) SetPluginAPI() *Lua {
	lua.L.SetGlobal("say", lua.L.NewFunction(lua.say))
	lua.L.SetGlobal("on", lua.L.NewFunction(lua.on))
	return lua
}

func (lua *Lua) say(L *LUA.LState) int {
	channel := L.ToString(1)
	message := L.ToString(2)
	lua.B.irc.Privmsg(channel, message)

	return 1
}

func (lua *Lua) on(L *LUA.LState) int {
	event := L.ToString(1)
	cback := L.ToString(2)

	lua.B.irc.AddCallback(event, func(event *irc.Event) {
		L.DoString(fmt.Sprintf(`%s("%s", "%s")`,
			cback,
			event.Arguments[0],
			event.Message()))
	})

	return 1
}

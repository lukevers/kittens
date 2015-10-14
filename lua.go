/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"github.com/yuin/gopher-lua"
)

type Lua struct {
	Lua     *lua.LState
	Bot     *Bot
	Channel *Channel
	events  map[string][]string
}

func NewLuaState(bot *Bot, channel *Channel) *Lua {
	L := &Lua{
		Lua:     lua.NewState(),
		Bot:     bot,
		Channel: channel,
		events:  make(map[string][]string),
	}

	defer L.Lua.Close()
	return L.SetPluginAPI()
}

func (L *Lua) SetPluginAPI() *Lua {
	L.Lua.SetGlobal("say", L.Lua.NewFunction(L.say))
	L.Lua.SetGlobal("on", L.Lua.NewFunction(L.on))
	L.Lua.SetGlobal("reload", L.Lua.NewFunction(L.reload))
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

	id := L.Bot.irc.AddCallback(event, func(event *irc.Event) {
		state.DoString(fmt.Sprintf(`%s("%s", "%s")`,
			cback,
			event.Arguments[0],
			event.Message()))
	})

	L.events[event] = append(L.events[event], id)
	return 1
}

func (L *Lua) reload(state *lua.LState) int {
	for event, ids := range L.events {
		for _, id := range ids {
			L.Bot.irc.RemoveCallback(event, id)
		}
	}

	L.Channel.LoadPlugins(L.Bot)
	return 1
}

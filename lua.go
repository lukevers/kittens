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
	Plugin  *Plugin
	eventId string
	eventType string
}

func NewLuaState(bot *Bot, channel *Channel, plugin *Plugin) *Lua {
	L := &Lua{
		Bot:     bot,
		Channel: channel,
		Plugin:  plugin,
		eventId:   "",
		eventType: "",
		Lua:     lua.NewState(lua.Options{
			IncludeGoStackTrace: true,
		}),
	}

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

	L.eventType = event
	L.eventId = L.Bot.irc.AddCallback(event, func(event *irc.Event) {
		if L.Channel.Name == event.Arguments[0] {

			state.DoString(fmt.Sprintf(`%s("%s", "%s")`,
				cback,
				event.Arguments[0],
				event.Message()))
		}
	})

	return 1
}

func (L *Lua) reload(state *lua.LState) int {
	for _, plugin := range L.Channel.Plugins {
		// Remove callback events
		L.Bot.irc.RemoveCallback(plugin.Lua.eventType, plugin.Lua.eventId)

		// Destroy lua
		defer plugin.Lua.Lua.Close()
		plugin.Lua = nil
	}

	L.Channel.LoadPlugins(L.Bot)
	return 1
}

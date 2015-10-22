/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/thoj/go-ircevent"
	"github.com/yuin/gopher-lua"
	"log"
)

type Lua struct {
	Lua       *lua.LState
	Bot       *Bot
	Channel   *Channel
	Plugin    *Plugin
	eventId   string
	eventType string
}

func NewLuaState(bot *Bot, channel *Channel, plugin *Plugin) *Lua {
	L := &Lua{
		Bot:       bot,
		Channel:   channel,
		Plugin:    plugin,
		eventId:   "",
		eventType: "",
		Lua: lua.NewState(lua.Options{
			IncludeGoStackTrace: true,
		}),
	}

	return L.SetPluginAPI()
}

func (L *Lua) SetPluginAPI() *Lua {
	L.Lua.SetGlobal("on", L.Lua.NewFunction(L.on))
	L.Lua.SetGlobal("reload", L.Lua.NewFunction(L.reload))
	L.Lua.SetGlobal("say", L.Lua.NewFunction(L.say))
	L.Lua.SetGlobal("join", L.Lua.NewFunction(L.join))
	L.Lua.SetGlobal("part", L.Lua.NewFunction(L.part))
	return L
}

func (L *Lua) on(state *lua.LState) int {
	event := state.ToString(1)
	cback := state.ToFunction(2)

	L.eventType = event
	L.eventId = L.Bot.irc.AddCallback(event, func(event *irc.Event) {
		if L.Channel.Name == event.Arguments[0] {
			table := new(lua.LTable)
			table.RawSetString("message", lua.LString(event.Message()))
			table.RawSetString("channel", lua.LString(event.Arguments[0]))
			table.RawSetString("nick", lua.LString(event.Nick))
			table.RawSetString("host", lua.LString(event.Host))
			table.RawSetString("source", lua.LString(event.Source))
			table.RawSetString("user", lua.LString(event.User))
			table.RawSetString("raw", lua.LString(event.Raw))

			if err := L.Lua.CallByParam(lua.P{
				Fn:      cback,
				NRet:    1,
				Protect: true,
			}, table); err != nil {
				log.Println("Error calling lua function: ", err)
			}
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

func (L *Lua) say(state *lua.LState) int {
	channel := state.ToString(1)
	message := state.ToString(2)
	L.Bot.irc.Privmsg(channel, message)

	return 1
}

func (L *Lua) join(state *lua.LState) int {
	channel := state.ToString(1)
	fresh := state.ToBool(2)

	// TODO - save everything/duplicate plugins (or not, depending), etc

	if fresh {
		// Do not duplicate plugin structure
	}

	L.Bot.irc.Join(channel)

	return 1
}

func (L *Lua) part(state *lua.LState) int {
	channel := state.ToString(1)
	hard := state.ToBool(2)

	// TODO - leave channel / set channel to disabled (or just destroy it, depending)

	if hard {
		// Completely remove from database
	}

	L.Bot.irc.Part(channel)

	return 1
}

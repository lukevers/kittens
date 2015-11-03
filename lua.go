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
	L.Channel.UnloadPlugins(L)
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

	var c *Channel
	var cc Channel
	db.Where(&Channel{
		Name:  channel,
		BotID: L.Bot.ID,
	}).Find(&cc)
	c = &cc

	if c.ID == 0 {
		if fresh {
			// New
		} else {
			// Clone
		}
	} else if !c.Enabled {
		// Check to see if already in bot slice
		exists := false
		for _, ch := range L.Bot.Channels {
			if ch.ID == c.ID {
				exists = true
				c = ch
				break
			}
		}

		if !exists {
			L.Bot.Channels = append(L.Bot.Channels, c)
		}

		c.Enable()
		L.Bot.irc.Join(channel)
		c.LoadPlugins(L.Bot)
	}

	return 1
}

func (L *Lua) part(state *lua.LState) int {
	channel := state.ToString(1)
	hard := state.ToBool(2)

	var c *Channel
	var cc Channel
	db.Where(&Channel{
		Name:  channel,
		BotID: L.Bot.ID,
	}).Find(&cc)
	c = &cc

	for _, ch := range L.Bot.Channels {
		if ch.ID == c.ID {
			c = ch
			break
		}
	}

	// Leave channel
	L.Bot.irc.Part(channel)
	c.UnloadPlugins(L)

	// Delete or disable channel and related
	if hard {
		c.Delete()
	} else {
		c.Disable()
	}

	return 1
}

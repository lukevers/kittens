/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"time"
)

type Channel struct {
	ID        int
	Name      string
	BotID     int `sql:"index"`
	Plugins   []*Plugin
	CreatedAt time.Time
	UpdatedAt time.Time
	Lua       *Lua
}

func (c *Channel) LoadPlugins(b *Bot) {
	c.Lua = NewLuaState(b, c)

	for _, plugin := range c.Plugins {
		if err := c.Lua.Lua.DoFile(plugin.Path); err != nil {
			panic(err)
		}
	}
}

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
	lua       *Lua
}

func (c *Channel) InitPlugins(b *Bot) {
	c.lua = NewLuaState(b)
	c.LoadPlugins()
}

func (c *Channel) LoadPlugins() {
	for _, plugin := range c.Plugins {
		if err := c.lua.Lua.DoFile(plugin.Path); err != nil {
			panic(err)
		}
	}
}

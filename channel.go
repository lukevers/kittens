/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"time"
)

type Channel struct {
	ID        int
	Name      string
	BotID     int `sql:"index"`
	Plugins   []Plugin
	CreatedAt time.Time
	UpdatedAt time.Time
	lua       *Lua
}

func (c *Channel) LoadPlugins(b *Bot) {
	c.lua = NewLuaState(b)

	// Add sample
	if err := c.lua.L.DoFile("plugins/echo.lua"); err != nil {
		panic(err)
	}
}

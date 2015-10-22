/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"log"
	"time"
)

type Channel struct {
	ID        int
	Name      string
	BotID     int `sql:"index"`
	Plugins   []*Plugin
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Channel) LoadPlugins(b *Bot) {
	for _, plugin := range c.Plugins {
		plugin.Lua = NewLuaState(b, c, plugin)

		if err := plugin.Lua.Lua.DoFile(plugin.Path); err != nil {
			log.Println("Erorr running plugin: ", err)
		}
	}
}

/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"github.com/yuin/gopher-lua"
	"time"
)

type Channel struct {
	ID        int
	Name      string
	BotID     int `sql:"index"`
	Plugins   []Plugin
	CreatedAt time.Time
	UpdatedAt time.Time
	lua       *lua.LState
}

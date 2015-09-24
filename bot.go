/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"time"
)

type Bot struct {
	ID          int
	Nickname    string
	Username    string
	Host        string
	Port        int
	UserID      int `sql:"index"`
	Channels    []Channel
	Enabled     bool `sql:"default:'0'"`
	DisplayName string
	lua         *lua.LState
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Channel struct {
	ID        int
	Name      string
	BotID     int `sql:"index"`
	Plugins   []Plugin
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetBot(by, value interface{}) *Bot {
	var bot Bot
	db.Where(fmt.Sprintf("%s = ?", by), value).First(&bot)
	return &bot
}

/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"github.com/yuin/gopher-lua"
	"time"
)

var bots map[int]*Bot = make(map[int]*Bot)

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
	CreatedAt   time.Time
	UpdatedAt   time.Time
	lua         *lua.LState
	bot         *irc.Connection
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

func (b *Bot) Connect() {
	b.Enabled = true
	db.Save(&b)

	// Create the IRC connection
	b.bot = irc.IRC(b.Nickname, b.Username)
	b.bot.Connect(fmt.Sprintf("%s:%d", b.Host, b.Port))

	// Join all channels
	for _, channel := range b.Channels {
		b.bot.Join(channel.Name)

		// TODO: Add plugins for each channel
	}
}

func (b *Bot) Disconnect() {
	b.bot.Quit()
	delete(bots, b.ID)

	b.Enabled = false
	db.Save(&b)
}

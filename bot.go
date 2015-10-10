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
	irc         *irc.Connection
}

func InitEnabledBots() {
	var enabledBots []*Bot
	db.Where(&Bot{Enabled: true}).Find(&enabledBots)
	for _, bot := range enabledBots {
		bot = bot.Related()

		bots[bot.ID] = bot
		bots[bot.ID].Connect()
	}
}

func GetBot(by, value interface{}) *Bot {
	var bot Bot
	db.Where(fmt.Sprintf("%s = ?", by), value).First(&bot)
	return &bot
}

func (b Bot) Related() *Bot {
	db.Model(&b).Related(&b.Channels)
	for c := range b.Channels {
		db.Model(&b.Channels[c]).Related(&b.Channels[c].Plugins)
	}

	return &b
}

func (b *Bot) Connect() {
	b.Enabled = true
	db.Save(&b)

	// Create the IRC connection
	b.irc = irc.IRC(b.Nickname, b.Username)
	b.irc.Connect(fmt.Sprintf("%s:%d", b.Host, b.Port))

	// Join all channels
	for _, channel := range b.Channels {
		// Setup Lua state for each channel
		channel.lua = lua.NewState()
		defer channel.lua.Close()

		go func(b *Bot, channel Channel) {
			channel.lua.SetGlobal("say", channel.lua.NewFunction(func(L *lua.LState) int {
				ch := L.ToString(1)
				msg := L.ToString(2)

				b.irc.Privmsg(ch, msg)

				return 1
			}))

			channel.lua.SetGlobal("on", channel.lua.NewFunction(func(L *lua.LState) int {
				event := L.ToString(1)
				cback := L.ToString(2)

				b.irc.AddCallback(event, func(event *irc.Event) {
					L.DoString(fmt.Sprintf(`%s("%s", "%s")`,
						cback,
						event.Arguments[0],
						event.Message()))
				})

				return 1
			}))
		}(b, channel)

		if err := channel.lua.DoFile("plugins/echo.lua"); err != nil {
			panic(err)
		}

		// Join all channels
		b.irc.Join(channel.Name)
	}
}

func (b *Bot) Disconnect() {
	b.irc.Quit()
	delete(bots, b.ID)

	b.Enabled = false
	db.Save(&b)
}

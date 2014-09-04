package main

import (
	"time"
)

type Bot struct {
	// ID is an int64 that is a bots identification
	// number.
	Id int64

	// Nick is a string with max-size set to 32 and
	// is the nickname that the bot will use on the
	// server it is connected to.
	Nick string `sql:"size:32"`

	// RealName is a string with max-size set to 255
	// and is the "real name" that the bot will use
	// on the server it is connected to.
	RealName string `sql:"size:255"`

	// Host is a string that defines
	Host string

	// Network is
	Network string

	// Port is
	Port int

	// Password is
	Password string

	// SSL is
	Ssl bool

	// Enabled is
	Enabled bool

	// UserID is a foreign key that references the
	// user that owns this bot.
	UserId int64

	// CreatedAt is a timestamp of when the specific
	// bot was created at.
	CreatedAt time.Time

	// UpdatedAt is a timestamp of when the specific
	// bot was last updated at.
	UpdatedAt time.Time
}

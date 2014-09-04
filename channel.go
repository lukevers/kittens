package main

import (
	"time"
)

type Channel struct {
	// ID is an int64 that is a channels identification
	// number.
	Id int64

	// Name is the name of the physical channel that
	// the bot on a specific server connects to.
	Name string

	// BotID is a foreign key that references the bot
	// that owns this channel.
	BotId int64

	// CreatedAt is a timestamp of when the specific channel
	// was created at.
	CreatedAt time.Time

	// UpdatedAt is a timestamp of when the specific channel
	// was last updated at.
	UpdatedAt time.Time
}

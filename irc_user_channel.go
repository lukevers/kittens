package main

import (
	"time"
)

type IrcUserChannel struct {
	// Id is a uint64 that is a channels identification number
	Id uint64

	// Channel
	Channel string

	// Modes
	Modes string

	// IrcUserId is a foreign key that references the irc user
	// that owns this channel.
	IrcUserId uint64

	// LastJoinedAt is a timestamp of when the irc user that
	// owns this channel last joined.
	LastJoinedAt time.Time

	// LastPartedAt is a timestamp of when the irc user that
	// owns this channel last parted/quit.
	LastPartedAt time.Time

	// CreatedAt is a timestamp of when the specific user's channel
	// was created at.
	CreatedAt time.Time

	// UpdatedAt is a timestamp of when a specific user's channel
	// was last updated at.
	UpdatedAt time.Time
}

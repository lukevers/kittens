package main

import (
	"time"
)

type IrcUser struct {
	// Id is a uint64 that is an irc users identification number
	Id uint64

	// Nickname is the name that the irc user goes by.
	Nickname string

	// Hostm is the host which the irc user identifies with.
	// The host is used with the nickname to 
	Host string

	// ServerId is a foreign key that references the server that
	// this user is on.
	ServerId uint64

	// CreatedAt is a timestamp of when the specific channel was
	// created at.
	CreatedAt time.Time

	// UpdatedAt is a timestamp of when the specific channel was
	// last updated at.
	UpdatedAt time.Time
}

/* vim: set autoindent noexpandtab tabstop=4 shiftwidth=4: */
package main

import (
	"time"
)

type Plugin struct {
	ID        int
	Name      string
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ChannelID int `sql:"index"`
}

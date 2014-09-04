package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

type Database struct {
	// Driver is a string that defines what type of database
	// we are using. As of right now we just support "mysql"
	// as driver, but should support others in the future.
	Driver string

	// Host is a string that defines the host of what we're
	// going to be connecting to. Generally it's "localhost"
	// or similar.
	Host string

	// Port is an integer that defines the port that we're
	// going to be connecting to. With MySQL it's generally
	// 3306.
	Port string

	// Name is the physical name of the database that
	// we're going to be connecting to. It really should be
	// "kittens" because that's what Kittens is called, but
	// it can be changed.
	Name string

	// Username is the username that's being used to connect
	// to the database.
	Username string

	// Password is the password that's being used to connect
	// to the database.
	Password string
}

//
func InitDatabase() {
	// Connect database
	db, err = sql.Open(config.Database.Driver,
		config.Database.Username+":"+
		config.Database.Password+"@tcp("+
		config.Database.Host+":"+
		config.Database.Port+")/"+
		config.Database.Name)
	if err != nil {
		warnf("Error connecting to database: %s", err)
		defer db.Close()
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		warnf("Error pinging database: %s", err)
		defer db.Close()
	}
}

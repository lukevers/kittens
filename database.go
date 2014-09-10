package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

var (
	db gorm.DB
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

// Init Database initializes the database, runs any migrations needed
// to be ran (with automigrate), and creates a default user if none
// exist.
func InitDatabase() {
	// Connect database
	//
	// mysql:
	// "username:password@tcp(host:port)/database?parseTime=true
	//
	db, err = gorm.Open(config.DB.Driver,
		config.DB.Username+":"+
			config.DB.Password+"@tcp("+
			config.DB.Host+":"+
			config.DB.Port+")/"+
			config.DB.Name+"?parseTime=true")
	if err != nil {
		warnf("Error connecting to database: %s", err)
		warn("Exiting with exit status 1")
		os.Exit(1)
	}

	// Test connection
	err = db.DB().Ping()
	if err != nil {
		warnf("Error pinging database: %s", err)
		warn("Exiting with exit status 1")
		os.Exit(1)
	}

	// Migrate/create tables
	verb("Running database auto migrate")

	//
	// Each child is connected to the parent via a foreign key
	// that relates to the parent's Id (which is a uint64).
	//
	// In this example below it's described as [Row of Table name]
	// (example name) where the [Row of Table name] is a row in
	// the table that is named, and (example name) is content that
	// could potentially be a field in one of the main columns in
	// that row. Here's an example of what it could look like:
	//
	// User (luke) 1:M
	//  │
	//  └─── Server (freenode) 1:M
	//        │
	//        ├─── Channel (#go-nuts) 1:1
	//        │
	//        ├─── Channel (#example) 1:1
	//        │
	//        ├─── Channel (#channel) 1:1
	//        │
	//        ├─── IrcUser (lukevers) 1:M
	//        │     │
	//        │     ├─── IrcUserChannel (#go-nuts) 1:1
	//        │     │
	//        │     └─── IrcUserChannel (#example) 1:1
	//        │
	//        └─── IrcUser (kittens) 1:M
	//              │
	//              ├─── IrcUserChannel (#example) 1:1
	//              │
	//              └─── IrcUserChannel (#channel) 1:1
	//

	db.AutoMigrate(User{}, Server{}, Channel{}, IrcUser{}, IrcUserChannel{})

	// Check to see if we have any users created.
	// If we don't have any users at all then we
	// need to make a default user.
	verb("Checking if any users exist")
	db.FirstOrCreate(&User{
		Username:  "admin",
		Password:  HashPassword("admin"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, &User{})
}

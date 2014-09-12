package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

var (
	db gorm.DB
)

// Init Database initializes the database, runs any migrations needed
// to be ran (with automigrate), and creates a default user if none
// exist.
func InitDatabase() {
	// Lowercase our driver flag to make it easier to parse
	*driverFlag = strings.ToLower(*driverFlag)

	// Check if driver is mysql
	if *driverFlag == "mysql" {
		// Check if ?parseTime=true is not included
		if !strings.Contains(*databaseFlag, "?parseTime=true") {
			// If `?parseTime=true` was not included then we need to
			// add it so MySQL works properly with Kittens.
			*databaseFlag += "?parseTime=true"
		}
	}

	// Check if driver flag is `sqlite`. It's a common typo to accidently
	// type `sqlite` instead of `sqlite3` (which we support), so to avoid
	// this completely, if we see `sqlite` anywhere in the driver string
	// we're just going to set it to `sqlite3` since that's the only
	// version of sqlite that we support.
	if strings.Contains(*driverFlag, "sqlite") {
		*driverFlag = "sqlite3"
	}

	// Check if driver flag contains `postgres`, and if it does then we just
	// want to change it to only be `postgres`. We're trying to avoid errors
	// in as many places as possible for the user.
	if strings.Contains(*driverFlag, "postgres") {
		*driverFlag = "postgres"
	}

	// Open connection
	db, err = gorm.Open(*driverFlag, *databaseFlag)
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
		Username: "admin",
		Password: HashPassword("admin"),
		Admin:    true,
		Twofa:    false,
	}, &User{})
}

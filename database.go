package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var (
	db gorm.DB
)

// Init Database initializes the database, runs any migrations needed
// to be ran (with automigrate), and creates a default user if none
// exist.
func InitDatabase() {
	// Figure out connection string that matches driver
	var conn string
	switch os.Getenv("DB_DRIVER") {
	case "sqlite3":
		conn = os.Getenv("DB_RESOURCE")
	case "mysql":
		conn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"))
	case "postgres":
		conn = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_SSLMODE"))
	}

	// Open connection
	var err error
	db, err = gorm.Open(os.Getenv("DB_DRIVER"), conn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// If we're in debug mode, debug database.
	if os.Getenv("APP_DEBUG") == "true" {
		db.LogMode(true)
	}

	// Test connection
	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	// Run database migrations
	log.Println("Running database migrations (if any)")
	db.AutoMigrate(&User{})

	/*
			// Create default user if no users exist
			var user User
			db.FirstOrCreate(&user, &User{
				Username: "default",
				Twofa:    false,
			})

		    db.First(&User{})

			if user.Password == "" {
				log.Println("Setting password for default user as secret")
				user.SetPassword("secret")
			}
	*/
}

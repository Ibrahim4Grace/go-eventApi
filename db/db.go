package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	// Check if database file exists and is accessible api.db is my db name
	if _, err := os.Stat("api.db"); os.IsNotExist(err) {
		file, err := os.Create("api.db")
		if err != nil {
			panic(fmt.Sprintf("Could not create database file: %v", err))
		}
		file.Close()
	}

	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to database: %v", err))
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("Could not ping database: %v", err))
	}
	///seting how many conn can can be open. if we av more than 10, rest will wait
	DB.SetMaxOpenConns(10)
	//seting how many conn we want to keep open if no one is using the db at the momemnt
	DB.SetMaxIdleConns(5)
	//making sure create table exists
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE NOT NULL,
		password TEXT NOT NULL
		)
		`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create users table: %v", err))
	}

	fmt.Println("Users database tables created successfully")

	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
            user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
        );
    `
	// the foreign key will ref user id from the users table

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create events table: %v", err))
	}

	fmt.Println("Events database tables created successfully")
}

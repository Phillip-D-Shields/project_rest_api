package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to db")
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
		`

	_, err := Db.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
		`

	_, err = Db.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
	}
}

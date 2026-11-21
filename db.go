package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite", "./linkscutter.db")
	fmt.Println("s")
	if err != nil {
		fmt.Println("1")
	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS links(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			original_url TEXT NOT NULL,
			mod_url TEXT NOT NULL UNIQUE

		)
	`
	_, err = db.Exec(createTableQuery)

	if err != nil {
		fmt.Println(err)
	}

	return db
}

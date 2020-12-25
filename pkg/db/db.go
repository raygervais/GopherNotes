package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	connection *sql.DB
}

func CreateDatabaseConnection(path string) Database {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		fmt.Printf("Error opening database connection: %s\n", err)
	}

	return Database{
		connection: db,
	}
}

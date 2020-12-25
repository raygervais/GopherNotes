package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raygervais/gophernotes/pkg/conf"
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

func (db Database) InitializeNotesTable() error {
	query := "CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, note TEXT, date TEXT)"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("Failed to create table: %s", err)
	}

	return nil
}

func (db Database) Create(message string) error {
	query := "INSERT INTO NOTES (note, date) VALUES (?, ?)"

	if len(message) == 0 {
		return fmt.Errorf("Invalid input provided as message parameter")
	}

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(message, time.Now().Format(conf.LayoutISO))
	if err != nil {
		return fmt.Errorf("Failed to create new note entry: %s", err)
	}

	return nil
}

func (db Database) Fetch() (*sql.Rows, error) {
	query := "SELECT note, date FROM notes"

	res, err := db.execQueryStatement(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db Database) Search(entry string) (*sql.Rows, error) {
	query := "SELECT note, date FROM notes WHERE note LIKE ? OR date LIKE ?"

	fmt.Println(query, entry)
	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Query(entry, entry)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (db Database) IterateOnRows(rows *sql.Rows) (string, error) {
	var output string

	var note, date string
	for rows.Next() {
		if err := rows.Scan(&note, &date); err != nil {
			return "", err
		}
		output += fmt.Sprintf("%v: %v\n", date, note)
	}

	return output, nil
}

// Helper Functions

func (db Database) prepareQueryStatement(query string) (*sql.Stmt, error) {
	return db.connection.Prepare(query)
}

func (db Database) execQueryStatement(query string) (*sql.Rows, error) {
	return db.connection.Query(query)
}

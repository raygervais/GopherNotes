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
	query := `CREATE VIRTUAL TABLE IF NOT EXISTS notes USING fts4 (
		note TEXT NOT NULL, 
		date TEXT NOT NULL)`

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
	query := "SELECT rowid, note, date FROM notes"

	return db.execQueryStatement(query)
}

func (db Database) Search(entry string) (*sql.Rows, error) {
	query := "SELECT rowid, note, date FROM notes WHERE note MATCH ?"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return nil, err
	}

	return stmt.Query(entry)
}

func (db Database) RetrieveByID(id int) (*sql.Row, error) {
	query := "SELECT note, date from notes where rowid = ?"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return nil, err
	}

	return stmt.QueryRow(id), nil
}

func (db Database) EditByID(id int, changes string) error {
	query := "UPDATE notes SET note = ? WHERE rowid = ?"

	if len(changes) == 0 {
		return fmt.Errorf("Invalid input provided as change parameter")
	}

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(changes, id)
	if err != nil {
		return fmt.Errorf("Failed to Edit note with rowid: %d\n%s", err)
	}
	return nil
}

func (db Database) DeleteByID(id int) error {
	query := "DELETE FROM notes WHERE rowid = ?"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("Failed to delete note with rowid: %d\n%s", id, err)
	}
	return nil
}

func (db Database) IterateOnRows(rows *sql.Rows) (string, error) {
	var output string

	var note, date, id string
	for rows.Next() {
		if err := rows.Scan(&id, &note, &date); err != nil {
			return "", err
		}
		output += fmt.Sprintf("%v) %v: %v\n", id, date, note)
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

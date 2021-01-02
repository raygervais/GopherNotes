package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" //Used as SQL driver
	"github.com/raygervais/gophernotes/pkg/conf"
)

// Database wraps around *sql.DB connection, imported by CommandLineInterface.
type Database struct {
	connection *sql.DB
}

// CreateDatabaseConnection creates or connects to an existing SQLite3 DB,
// path provided must be absolute path to file.
func CreateDatabaseConnection(path string) Database {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		fmt.Printf("Error opening database connection: %s\n", err)
	}

	return Database{
		connection: db,
	}
}

// InitializeNotesTable creates the virtual NOTES table,
// leveraging fts4 for whole-world pattern matching.
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

// Create enters new note entry into DB with current timestamp.
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

// Fetch retrieves all notes from DB based on order of entry.
func (db Database) Fetch() (*sql.Rows, error) {
	query := "SELECT rowid, note, date FROM notes"

	return db.execQueryStatement(query)
}

// Search retrieves all notes which comply with the pattern match provided.
func (db Database) Search(entry string) (*sql.Rows, error) {
	query := "SELECT rowid, note, date FROM notes WHERE note MATCH ?"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return nil, err
	}

	return stmt.Query(entry)
}

// RetrieveByID finds and returns a single record which complies with the same rowid provided.
func (db Database) RetrieveByID(id int) (*sql.Row, error) {
	query := "SELECT note, date from notes where rowid = ?"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return nil, err
	}

	return stmt.QueryRow(id), nil
}

// EditByID processes updating the note entry with the corresponding rowid with changes supplied.
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
		return fmt.Errorf("Failed to Edit note with rowid: %d\n%s", id, err)
	}
	return nil
}

// DeleteByID processes deletion of a single entry with the supplied rowid.
func (db Database) DeleteByID(id int) error {
	query := "DELETE FROM notes WHERE rowid = ?"

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("Failed to delete note with rowid: %d\n%s", id, err)
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return fmt.Errorf("0 rows were deleted with rowid: %d", id)
	}

	return nil
}

// IterateOnRows is a helper which formats SQL Row values to preferred configuration.
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

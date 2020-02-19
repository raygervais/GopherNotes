package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

/* 
 * 
 * Database Helpers
 *
 */

func PrepareCommandStatement(database *sql.DB, command string) (*sql.Stmt, error) {
	return database.Prepare(command)
}

func QueryCommandStatement(database *sql.DB, command string) (*sql.Rows, error) {
	return database.Query(command)
}

func ExecuteStatementHandler(statement *sql.Stmt, arguments ...interface{}) sql.Result {
	result, err := statement.Exec(arguments...)
	ErrorHandler(err)
	return result
}

func DatabaseStatementErrorHandler(value *sql.Stmt, err error) *sql.Stmt {
	ErrorHandler(err)
	return value
}

func DatabaseRowsErrorHandler(rows *sql.Rows, err error) *sql.Rows {
	ErrorHandler(err)

	return rows
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

/* 
 * 
 * Database Actions
 *
 */
// CreateDatabase - Creates Database Connection to Given SQLite DB Path
func CreateDatabaseConnection(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	ErrorHandler(err)

	return db
}

// InitializeNotesTable - Creates Notes Table with Provided Database Connection
func InitializeNotesTable(database *sql.DB) sql.Result {
	return ExecuteStatementHandler(
		DatabaseStatementErrorHandler(
			PrepareCommandStatement(database, "CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, entry TEXT, date TEXT)")))

}

// InsertIntoNotesTable - Inserts New Note Into Table
func InsertIntoNotesTable(database *sql.DB, note Note) sql.Result {
	return ExecuteStatementHandler(
		DatabaseStatementErrorHandler(
			PrepareCommandStatement(database, "INSERT INTO notes (entry, date) VALUES (?, ?)")), note.entry, note.date)
}

// RetrieveNotes - Gets All Notes from Database
func RetrieveNotes(database *sql.DB) *sql.Rows {
	return DatabaseRowsErrorHandler(
		QueryCommandStatement(
			database, "Select id, entry, date from notes"))
}

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

const (
	panicError   = "this should cause a panic"
	testLocation = "./test.db"
)

/*
 *
 * Testing Helpers
 *
 */
func RecoverFromPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}

func DeleteDatabase(t *testing.T, path string) {
	ErrorHandler(os.Remove(testLocation))
}

func PrintTransactionStatus(t *testing.T, transaction sql.Result) {
	rows, _ := transaction.RowsAffected()
	lastID, _ := transaction.LastInsertId()
	fmt.Printf("Rows Affected: %v, Last ID: %v\n", rows, lastID)
}

/*
 *
 * Tests
 *
 */
func TestDatabaseFunctions(t *testing.T) {
	t.Run("Should not error out if err is nil", func(t *testing.T) {
		ErrorHandler(nil)
	})

	t.Run("Should error out when err is not nil", func(t *testing.T) {
		defer RecoverFromPanic(t)

		ErrorHandler(errors.New(panicError))
	})

	t.Run("Should create database.db if it doesn't exist", func(t *testing.T) {
		CreateDatabaseConnection(testLocation)
	})

	t.Run("Should create table if it doesn't exist", func(t *testing.T) {
		db := CreateDatabaseConnection(testLocation)
		InitializeNotesTable(db)

		DeleteDatabase(t, testLocation)

	})

	t.Run("Should insert into table a note entry", func(t *testing.T) {
		db := CreateDatabaseConnection(testLocation)
		InitializeNotesTable(db)
		InsertIntoNotesTable(db, Note{
			entry: "Hello, Testing World!",
			date:  time.Now().Format(LayoutISO),
		})
	})

	t.Run("Should retrieve from table a note specified entry", func(t *testing.T) {
		db := CreateDatabaseConnection(testLocation)
		rows := SearchNotesByDate(db, time.Now().Format(LayoutISO))

		if !rows.Next() {
			t.Fatal("Should have returned an entry")
		}

		DeleteDatabase(t, testLocation)
	})
}

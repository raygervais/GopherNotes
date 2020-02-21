package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	NotesLocation = "./gophernotes.db"
	LayoutISO     = "2006-01-02"
)

// Application Entry Point
func main() {
	// Create database if it doesn't exist
	// Then connection for further use
	database := CreateDatabaseConnection(NotesLocation)
	InitializeNotesTable(database)

	// Declare CLI Options
	newCommand := GenerateFlag("new")
	searchCommand := GenerateFlag("search")

	// New Flag
	newTextPtr := GenerateFlagParams(newCommand, "text", "", "note text to be saved")

	// Search Flag
	searchTextPtr := GenerateFlagParams(searchCommand, "text", "", "note text to be searched")
	searchDatePtr := GenerateFlagParams(searchCommand, "date", "", "note date to be searched")

	// Confirm more than two args provided
	VerifyFlagArguments(os.Args)

	// Subcommand parsing
	switch os.Args[1] {
	case "new":
		newCommand.Parse(os.Args[2:])
	case "search":
		searchCommand.Parse(os.Args[2:])
	default:
		fmt.Println("Subargument Parsing failed!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Validate Required Fields
	if newCommand.Parsed() {
		// Required Flags
		VerifyFlagParams(newTextPtr, newCommand)

		note := Note{
			entry: *newTextPtr,
			date:  time.Now().Format(LayoutISO),
		}

		InsertIntoNotesTable(database, note)
		rows := RetrieveNotes(database)

		var entry, date string
		var id int
		for rows.Next() {
			rows.Scan(&id, &entry, &date)
			fmt.Printf("%v: %v\n", date, entry)
		}
	}

	if searchCommand.Parsed() {
		// Search by Text
		if *searchTextPtr != "" && *searchDatePtr == "" {

			notes := ParseDatabaseRows(
				SearchNotesByEntry(database, *searchTextPtr))

			fmt.Println(PrintNoteOutput(notes))

			// Search by Date
		} else if *searchDatePtr != "" && *searchTextPtr == "" {
			notes := ParseDatabaseRows(
				SearchNotesByDate(database, *searchDatePtr))

			fmt.Println(PrintNoteOutput(notes))

		} else {
			fmt.Println("Search Parsing failed!")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}
}

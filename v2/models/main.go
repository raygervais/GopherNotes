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

// Debugging
var NOTES = []Note{
	{date: "2020-02-13", entry: "Hello Ray!"},
	{date: "2020-02-14", entry: "Hello, World!"},
}

// Application Entry Point
func main() {
	// Subcommand Parents
	newCommand := flag.NewFlagSet("new", flag.ExitOnError)
	searchCommand := flag.NewFlagSet("search", flag.ExitOnError)

	// Create Database if it doesn't exist, then Connection
	database := CreateDatabaseConnection(NotesLocation)
	InitializeNotesTable(database)

	// New Note Command Flag Pointers
	newTextPtr := newCommand.String("text", "", "Note text to be saved")

	// Search Note Command Flag Pointers
	searchTextPtr := searchCommand.String("text", "", "Search notes with the following text")

	// Verify we are providing a subcomamnd
	// os.Arg[0] is the main command
	// os.Arg[1] is the subcommand
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command, new or search")
		os.Exit(1)
	}

	// Switch on subcommand parsing
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
		if *newTextPtr == "" {
			newCommand.PrintDefaults()
			os.Exit(1)
		}

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
		// Required Flags
		if *searchTextPtr == "" {
			searchCommand.PrintDefaults()
			os.Exit(1)
		}

		// 	// Call Search Note Controller
		// 	notes := CreateSearchResults(NOTES, *searchTextPtr)
		// 	if len(notes) == 0 {
		// 		fmt.Println("Failed to find:", *searchTextPtr)
		// 		os.Exit(1)
		// 	}

		// 	fmt.Println("Search Results:")
		// 	fmt.Println(notes)
	}
}
